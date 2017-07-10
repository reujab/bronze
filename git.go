package main

import (
	"bufio"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	// only run git commands if in a git repo
	dir, err := os.Getwd()
	check(err)
	dir, err = filepath.EvalSymlinks(dir)
	check(err)
	var isRepo bool
	for dir != "/" {
		file, err := os.Stat(filepath.Join(dir, ".git"))
		if err == nil && file.IsDir() {
			isRepo = true
			break
		}
		dir = filepath.Dir(dir)
	}
	if !isRepo {
		return
	}

	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(5)
	var domain string
	var stashes int
	var commitsAhead int
	var branch string
	var dirty, modified, staged bool

	go func() {
		defer waitgroup.Done()

		cmd := exec.Command("git", "status", "--porcelain")
		stdout, err := cmd.StdoutPipe()
		check(err)
		check(cmd.Start())

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			dirty = true

			// https://git-scm.com/docs/git-status
			x := scanner.Text()[0]
			y := scanner.Text()[1]
			if x == ' ' && (y == 'M' || y == 'D') {
				modified = true
			} else if x != '?' {
				staged = true
			}
		}
		check(scanner.Err())
		check(cmd.Wait())
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

		check(cmd.Wait())
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

		cmd.Wait()
	}()

	go func() {
		defer waitgroup.Done()

		stdout, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			// no commits yet
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
			// no configured remote
			return
		}
		uri, err := url.Parse(string(stdout))
		check(err)
		domain = uri.Hostname()
	}()

	waitgroup.Wait()

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
	if stashes != 0 || commitsAhead != 0 {
		segments = append(segments, strings.Repeat(iconStash, stashes)+strings.Repeat(iconAhead, commitsAhead))
	}
	if branch != "" {
		segments = append(segments, branch)
	}
	if dirty {
		segment.background = "yellow"

		var icons string
		if modified {
			icons += iconCircle
		}
		if staged {
			icons += iconPlus
		}
		if icons != "" {
			segments = append(segments, icons)
		}
	}
	segment.visible = true
	segment.value = strings.Join(segments, " ")
}
