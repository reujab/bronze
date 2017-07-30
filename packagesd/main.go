package main

import (
	"bufio"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var packages int // packages to update

func main() {
	sock, err := net.Listen("unix", "/tmp/packagesd.sock")
	check(err)
	defer func() { check(sock.Close()) }()
	check(os.Chmod("/tmp/packagesd.sock", 0666))

	getPackages()
	go func() {
		for {
			time.Sleep(time.Minute * 10)
			getPackages()
		}
	}()

	go func() {
		for {
			conn, err := sock.Accept()
			check(err)
			_, _ = conn.Write([]byte(strconv.Itoa(packages)))
			check(conn.Close())
		}
	}()

	// gracefully exit on interrupt, closing the socket
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getPackages() {
	packages = 0

	_, err := exec.LookPath("pacaur")
	if err == nil {
		queryPacaur()
		return
	}

	_, err = exec.LookPath("dnf")
	if err == nil {
		queryDnf()
		return
	}

	_, err = exec.LookPath("apt")
	if err == nil {
		queryApt()
		return
	}
}

func queryPacaur() {
	// update pacman databases
	cmd := exec.Command("pacman", "-Syy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	// count packages with available updates
	cmd := exec.Command("pacaur", "-Qu")
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		packages++
	}
	check(scanner.Err())
	_ = cmd.Wait()
}

func queryDnf() {
	// count packages with available updates
	cmd := exec.Command("dnf", "check-update", "-q")
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		packages++
	}
	// account for the blank line
	packages--
	check(scanner.Err())
	_ = cmd.Wait()
}

func queryApt() {
	// update apt databases
	check(exec.Command("apt", "update").Run())

	// count packages with available updates
	cmd := exec.Command("apt", "list", "--upgradable")
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		packages++
	}
	// account for the first line, "Listing... Done"
	packages--
	check(scanner.Err())
	_ = cmd.Wait()
}
