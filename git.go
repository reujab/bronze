package main

import (
	"net/url"
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
	defer repo.Free()

	var domain string
	remote, err := repo.Remotes.Lookup("origin")
	if err == nil {
		uri, err := url.Parse(remote.Url())
		check(err)
		domain = uri.Hostname()
		remote.Free()
	}

	var stashes int
	repo.Stashes.Foreach(func(int, string, *git.Oid) error {
		stashes++
		return nil
	})

	var commitsAhead int
	head, err := repo.Head()
	check(err)
	upstream, err := head.Branch().Upstream()
	if err == nil {
		commitsAhead, _, err = repo.AheadBehind(head.Branch().Reference.Target(), upstream.Target())
		check(err)
	}

	var branch string
	branch, err = head.Branch().Name()
	check(err)
	head.Free()

	var dirty, modified, staged bool
	status, err := repo.StatusList(&git.StatusOptions{
		Flags: git.StatusOptIncludeUntracked,
	})
	check(err)
	count, err := status.EntryCount()
	check(err)
	if count != 0 {
		dirty = true
	}
	for i := 0; i < count; i++ {
		entry, err := status.ByIndex(i)
		check(err)
		if entry.Status&git.StatusWtNew != 0 || entry.Status&git.StatusWtModified != 0 || entry.Status&git.StatusWtDeleted != 0 || entry.Status&git.StatusWtTypeChange != 0 || entry.Status&git.StatusWtRenamed != 0 {
			modified = true
		}
		if entry.Status&git.StatusIndexNew != 0 || entry.Status&git.StatusIndexModified != 0 || entry.Status&git.StatusIndexDeleted != 0 || entry.Status&git.StatusIndexRenamed != 0 || entry.Status&git.StatusIndexTypeChange != 0 {
			staged = true
		}
	}
	status.Free()

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
