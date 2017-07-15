package main

import (
	"os"
	"os/user"
)

func userSegment(segment *segment) {
	usr, err := user.Current()
	check(err)
	hostname, err := os.Hostname()
	check(err)
	segment.value = usr.Username + "@" + hostname
}
