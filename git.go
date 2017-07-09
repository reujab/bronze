package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

// the git segment provides useful information about a git repository such as the domain of the "origin" remote (with an icon), the current branch, and whether the HEAD is dirty
func gitSegment(segment *segment) {
	dir, err := os.Getwd()
	check(err)
	dir, err = filepath.EvalSymlinks(dir)
	check(err)

	// go through every parent directory looking for a .git directory
	for dir != "/" {
		file, err := os.Stat(filepath.Join(dir, ".git"))
		if err != nil || !file.IsDir() {
			dir = filepath.Dir(dir)
			continue
		}

		// found a git repository
		config, err := ini.Load(filepath.Join(dir, ".git/config"))
		if err != nil {
			return
		}
		origin, err := url.Parse(config.Section(`remote "origin"`).Key("url").Value())
		if err != nil {
			return
		}

		fmt.Println(origin.Hostname())

		break
	}
}
