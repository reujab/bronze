package main

import (
	"bufio"
	"net/url"
	"os/exec"
	"strings"
	"sync"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(2)
	var isRepo bool
	var domain string
	var modified, indexModified bool

	go func() {
		defer waitgroup.Done()

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
	}()

	go func() {
		defer waitgroup.Done()

		stdout, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
		if err != nil {
			return
		}

		uri, err := url.Parse(string(stdout))
		check(err)
		domain = uri.Hostname()
	}()

	waitgroup.Wait()
	if !isRepo {
		return
	}

	var segments []string
	switch domain {
	case "github.com":
		segments = append(segments, iconGithub)
	case "gitlab.com":
		segments = append(segments, iconGitlab)
	case "bitbucket.com":
		segments = append(segments, iconBitbucket)
	default:
		segments = append(segments, iconGit)
	}
	if modified || indexModified {
		segment.background = "yellow"

		var icons string
		if modified {
			icons += iconModified
		}
		if indexModified {
			icons += iconIndexModified
		}
		segments = append(segments, icons)
	}
	segment.visible = true
	segment.value = strings.Join(segments, " ")
}
