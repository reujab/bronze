package main

import (
	"os"
	"os/user"

	. "github.com/reujab/bronze/types"
)

func userSegment(segment *Segment) {
	usr, err := user.Current()
	die(err)
	hostname, err := os.Hostname()
	die(err)
	segment.Value = usr.Username + "@" + hostname
}
