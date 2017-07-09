package main

import (
	"os/exec"
	"strings"
)

func nodeSegment(segment *segment) {
	output, err := exec.Command("node", "--version").Output()
	check(err)
	segment.visible = true
	segment.value = strings.TrimSpace(string(output))
}
