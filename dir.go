package main

import (
	"os"
	"os/user"
	"strconv"
	"strings"

	. "github.com/reujab/bronze/types"
)

func dirSegment(segment *Segment) {
	usr, err := user.Current()
	die(err)
	segment.Value, err = os.Getwd()
	die(err)
	length, _ := strconv.Atoi(os.Getenv("BRONZE_DIR_LENGTH"))

	if strings.HasPrefix(segment.Value, usr.HomeDir) {
		segment.Value = icons["home"] + segment.Value[len(usr.HomeDir):]
	}

	if length != 0 {
		split := strings.Split(segment.Value, string(os.PathSeparator))
		for i := len(split) - 2; i >= 0; i-- {
			runes := []rune(split[i])
			if len(runes) > length {
				split[i] = string(runes[:length])
			}
		}

		segment.Value = strings.Join(split, string(os.PathSeparator))
	}
}
