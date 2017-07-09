package main

import (
	"bufio"
	"os/exec"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	cmd := exec.Command("git", "status", "--porcelain")
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	scanner := bufio.NewScanner(stdout)
	var modified, indexModified bool
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == 'M' || line[1] == 'M' {
			modified = true
		} else if line[0] == 'A' || line[1] == 'A' || line[0] == 'D' {
			indexModified = true
		}
	}
	check(scanner.Err())

	if cmd.Wait() != nil {
		return
	}

	segment.visible = true
	if modified || indexModified {
		segment.background = "yellow"
		if modified {
			segment.value += iconModified
		}
		if indexModified {
			segment.value += iconIndexModified
		}
	}
}
