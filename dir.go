package main

import (
	"os"
	"os/user"
	"strings"
)

func dirSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	cwd, err := os.Getwd()
	check(err)

	if strings.HasPrefix(cwd, usr.HomeDir) {
		segment.value = icons["home"] + cwd[len(usr.HomeDir):]
	} else {
		segment.value = cwd
	}
}
