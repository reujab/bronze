package main

import (
	"os"
	"strings"

	"github.com/libgit2/git2go"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	dir, err := os.Getwd()
	check(err)
	repo, err := git.OpenRepositoryExtended(dir, 0, "/")
	if err != nil {
		return
	}

	var domain string
	var stashes int
	var commitsAhead int
	var branch string
	var dirty, modified, staged bool

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
