package main

import (
	"os"
	"os/user"
	"strconv"
	"strings"
	"unicode/utf8"
)

func dirSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	segment.value, err = os.Getwd()
	check(err)
	length, _ := strconv.Atoi(os.Getenv("BRONZE_DIR_LENGTH"))

	if strings.HasPrefix(segment.value, usr.HomeDir) {
		segment.value = icons["home"] + segment.value[len(usr.HomeDir):]
	}

	if length != 0 {
		split := strings.Split(segment.value, string(os.PathSeparator))
		for i := len(split) - 2; i >= 0; i-- {
			char, size := utf8.DecodeRuneInString(split[i])
			if size != 0 {
				split[i] = string(char)
			}
		}

		segment.value = strings.Join(split, string(os.PathSeparator))
	}
}
