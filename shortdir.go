package main

import (
	"os"
	"os/user"
	"strings"
	"unicode/utf8"
)

func shortDirSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	cwd, err := os.Getwd()
	check(err)

	if strings.HasPrefix(cwd, usr.HomeDir) {
		cwd = icons["home"] + cwd[len(usr.HomeDir):]
	}

	split := strings.Split(cwd, string(os.PathSeparator))
	for i := len(split) - 2; i >= 0; i-- {
		char, size := utf8.DecodeRuneInString(split[i])
		if size != 0 {
			split[i] = string(char)
		}
	}

	segment.value = strings.Join(split, string(os.PathSeparator))
}
