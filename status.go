package main

import (
	"os"
	"os/user"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func statusSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	if usr.Uid == "0" {
		segment.value += icons["root"]
	}

	if unix.Access(".", unix.W_OK) != nil {
		segment.value += icons["readonly"]
	}

	status, err := strconv.Atoi(os.Getenv("STATUS"))
	check(err)
	if status != 0 {
		segment.value += icons["failed"]
	}

	jobs, err := strconv.Atoi(os.Getenv("JOBS"))
	check(err)
	segment.value += strings.Repeat(icons["job"], jobs)
}
