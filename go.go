package main

import (
	"os/exec"
	"strings"
)

func goSegment(segment *segment) {
	stdout, err := exec.Command("go", "version").Output()
	check(err)
	segment.visible = true
	segment.value = strings.Split(string(stdout), " ")[2]
}
