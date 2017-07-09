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
		segment.value += iconLightning
	}

	if unix.Access(".", unix.W_OK) != nil {
		segment.value += iconLock
	}

	status, err := strconv.Atoi(os.Getenv("STATUS"))
	check(err)
	if status != 0 {
		segment.value += iconExclam
	}

	jobs, err := strconv.Atoi(os.Getenv("JOBS"))
	check(err)
	segment.value += strings.Repeat(iconGear, jobs)

	segment.visible = segment.value != ""
}
