package main

import (
	"os"
	"os/user"

	. "github.com/reujab/bronze/types"
)

func userSegment(segment *Segment) {
	usr, err := user.Current()
	check(err)
	hostname, err := os.Hostname()
	check(err)
	segment.Value = usr.Username + "@" + hostname
}
