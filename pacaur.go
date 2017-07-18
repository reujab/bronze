package main

import (
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	. "github.com/reujab/bronze/types"
)

func pacaurSegment(segment *Segment) {
	conn, err := net.Dial("unix", "/tmp/pacaurd.sock")
	check(err)
	defer func() { check(conn.Close()) }()

	packages, err := ioutil.ReadAll(conn)
	check(err)
	num, err := strconv.Atoi(string(packages))
	check(err)

	if num > 5 {
		segment.Value = icons["package"] + string(packages)
	} else {
		segment.Value = strings.Repeat(icons["package"], num)
	}
}
