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

	segment.visible = true
	if strings.HasPrefix(cwd, usr.HomeDir) {
		segment.value = "\uf015" + cwd[len(usr.HomeDir):]
	} else {
		segment.value = cwd
	}
}
