package main

import (
	"encoding/json"
	"net"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/reujab/RSSd/commands"
	. "github.com/reujab/bronze/types"
)

func rssSegment(segment *Segment) {
	usr, err := user.Current()
	die(err)
	conn, err := net.Dial("unix", filepath.Join(usr.HomeDir, ".local/share/rssd.sock"))
	die(err)

	conn.Write([]byte{commands.List})
	var items []interface{}
	die(json.NewDecoder(conn).Decode(&items))

	if len(items) > 5 {
		segment.Value = icons["rss"] + strconv.Itoa(len(items))
	} else {
		segment.Value = strings.Repeat(icons["rss"], len(items))
	}
}
