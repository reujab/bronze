package main

import (
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	. "github.com/reujab/bronze/types"
)

func packagesSegment(segment *Segment) {
	conn, err := net.Dial("unix", "/tmp/packagesd.sock")
	die(err)
	defer func() { die(conn.Close()) }()

	packages, err := ioutil.ReadAll(conn)
	die(err)
	num, err := strconv.Atoi(string(packages))
	die(err)

	if num > 5 {
		segment.Value = icons["package"] + string(packages)
	} else {
		segment.Value = strings.Repeat(icons["package"], num)
	}
}
