package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	. "github.com/reujab/bronze/types"
	"gopkg.in/libgit2/git2go.v26"
)

// Reformat scp-like url (e. g. ssh)
// https://github.com/motemen/ghq/blob/master/url.go
func fixUrl(url string) (string) {
	var hasSchemePattern = regexp.MustCompile("^[^:]+://")
	var scpLikeUrlPattern = regexp.MustCompile("^([^@]+@)?([^:]+):/?(.+)$")

	if !hasSchemePattern.MatchString(url) && scpLikeUrlPattern.MatchString(url) {
		matched := scpLikeUrlPattern.FindStringSubmatch(url)
		user := matched[1]
		host := matched[2]
		path := matched[3]

		url = fmt.Sprintf("ssh://%s%s/%s", user, host, path)
	}

	return url
}

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *Segment) {
	dir, err := os.Getwd()
	die(err)
	repo, err := git.OpenRepositoryExtended(dir, 0, "/")
	if err != nil {
		return
	}
	defer repo.Free()

	var domainName string
	remote, err := repo.Remotes.Lookup("origin")
	if err == nil {
		remoteUrl := fixUrl(remote.Url())
		uri, err := url.Parse(remoteUrl)
		if err == nil && len(uri.Hostname()) > 4 {
			// strip the tld off the hostname
			domainName = uri.Hostname()[:len(uri.Hostname())-4]
		}
		remote.Free()
	}

	var stashes int
	repo.Stashes.Foreach(func(int, string, *git.Oid) error {
		stashes++
		return nil
	})

	var ahead, behind int
	var branch string
	head, err := repo.Head()
	if err == nil {
		upstream, err := head.Branch().Upstream()
		if err == nil {
			ahead, behind, err = repo.AheadBehind(head.Branch().Target(), upstream.Target())
			die(err)
		}

		branch, err = head.Branch().Name()
		if err != nil {
			// head is detached
			branch = head.Branch().Target().String()[:7]
		}
		head.Free()
	}

	var dirty, modified, staged bool
	status, err := repo.StatusList(&git.StatusOptions{
		Flags: git.StatusOptIncludeUntracked,
	})
	if err != nil {
		// bare repository
		return
	}
	count, err := status.EntryCount()
	die(err)
	if count != 0 {
		dirty = true
	}
	for i := 0; i < count; i++ {
		entry, err := status.ByIndex(i)
		die(err)
		if entry.Status&git.StatusWtNew != 0 || entry.Status&git.StatusWtModified != 0 || entry.Status&git.StatusWtDeleted != 0 || entry.Status&git.StatusWtTypeChange != 0 || entry.Status&git.StatusWtRenamed != 0 {
			modified = true
		}
		if entry.Status&git.StatusIndexNew != 0 || entry.Status&git.StatusIndexModified != 0 || entry.Status&git.StatusIndexDeleted != 0 || entry.Status&git.StatusIndexRenamed != 0 || entry.Status&git.StatusIndexTypeChange != 0 {
			staged = true
		}
	}
	status.Free()

	var segments []string
	domainIcon := icons[domainName]
	if domainIcon == "" {
		domainIcon = icons["git"]
	}
	if domainIcon != "" {
		segments = append(segments, domainIcon)
	}
	if stashes != 0 || ahead != 0 || behind != 0 {
		section := strings.Repeat(icons["stash"], stashes) + strings.Repeat(icons["ahead"], ahead) + strings.Repeat(icons["behind"], behind)
		if section != "" {
			segments = append(segments, section)
		}
	}
	if branch != "" {
		segments = append(segments, branch)
	}
	if dirty {
		segment.Background = "yellow"

		var section string
		if modified {
			section += icons["modified"]
		}
		if staged {
			section += icons["staged"]
		}
		if section != "" {
			segments = append(segments, section)
		}
	}
	segment.Value = strings.Join(segments, " ")
}
