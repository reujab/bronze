package main

import (
	"os"
	"os/user"
	"strconv"
	"strings"

	. "github.com/reujab/bronze/types"
	"golang.org/x/sys/unix"
)

func statusSegment(segment *Segment) {
	usr, err := user.Current()
	check(err)
	if usr.Uid == "0" {
		segment.Value += icons["root"]
	}

	if unix.Access(".", unix.W_OK) != nil {
		segment.Value += icons["readonly"]
	}

	status, err := strconv.Atoi(os.Getenv("status"))
	check(err)
	if status != 0 {
		segment.Value += icons["failed"]
	}

	jobs, err := strconv.Atoi(os.Getenv("jobs"))
	check(err)
	segment.Value += strings.Repeat(icons["job"], jobs)
}
