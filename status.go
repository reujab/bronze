package main

import (
	"os"
	"os/user"
	"strconv"
)

func statusSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	if usr.Uid == "0" {
		segment.value += iconLightning
	}

	status, err := strconv.Atoi(os.Getenv("STATUS"))
	check(err)
	if status != 0 {
		segment.value += iconExclam
	}

	segment.visible = segment.value != ""
}
