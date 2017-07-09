package main

import (
	"bufio"
	"os/exec"
	"sync"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(1)
	var isRepo bool
	var modified, indexModified bool

	go func() {
		cmd := exec.Command("git", "status", "--porcelain")
		stdout, err := cmd.StdoutPipe()
		check(err)
		check(cmd.Start())

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			// https://git-scm.com/docs/git-status
			if line[0] == 'M' || line[1] == 'M' {
				modified = true
			} else if line[0] == 'A' || line[1] == 'A' || line[0] == 'D' {
				indexModified = true
			}
		}
		check(scanner.Err())

		// the directory is in a git repo if the `git status` command exited successfully
		isRepo = cmd.Wait() == nil
		waitgroup.Done()
	}()

	waitgroup.Wait()
	segment.visible = isRepo
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
