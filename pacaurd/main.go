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
	sock, err := net.Listen("unix", "/tmp/pacaurd.sock")
	check(err)
	defer func() { check(sock.Close()) }()

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
			conn.Write([]byte(strconv.Itoa(packages)))
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

	cmd := exec.Command("pacaur", "-Qu")
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		packages++
	}
	check(scanner.Err())
	cmd.Wait()
}
