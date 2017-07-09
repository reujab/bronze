package main

import (
	"bufio"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(5)
	var isRepo bool
	var domain string
	var stashes int
	var commitsAhead int
	var branch string
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

		cmd := exec.Command("git", "stash", "list")
		stdout, err := cmd.StdoutPipe()
		check(err)
		check(cmd.Start())

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			stashes++
		}
		check(scanner.Err())
	}()

	go func() {
		defer waitgroup.Done()

		cmd := exec.Command("git", "log", "--oneline", "@{upstream}..")
		stdout, err := cmd.StdoutPipe()
		check(err)
		check(cmd.Start())

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			commitsAhead++
		}
		check(scanner.Err())
	}()

	go func() {
		defer waitgroup.Done()

		stdout, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			return
		}
		branch = strings.TrimSpace(string(stdout))

		// if the head is detached, display the short hash
		if branch == "HEAD" {
			stdout, err = exec.Command("git", "rev-parse", "--short", "HEAD").Output()
			check(err)
			branch = strings.TrimSpace(string(stdout))
		}
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
	if stashes > 5 {
		segments = append(segments, iconStash+strconv.Itoa(stashes))
	} else if stashes != 0 {
		segments = append(segments, strings.Repeat(iconStash, stashes))
	}
	if commitsAhead > 5 {
		segments = append(segments, iconAhead+strconv.Itoa(commitsAhead))
	} else if commitsAhead != 0 {
		segments = append(segments, strings.Repeat(iconAhead, commitsAhead))
	}
	segments = append(segments, branch)
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
