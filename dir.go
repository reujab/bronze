package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/reujab/bronze/types"
)

func dirSegment(segment *Segment) {
	var aliases []string
	if os.Getenv("BRONZE_DIR_ALIASES") != "" {
		aliases = strings.Split(os.Getenv("BRONZE_DIR_ALIASES"), ":")
	}
	if len(aliases)%2 != 0 {
		dief("dir: invalid BRONZE_DIR_ALIASES")
	}

	usr, err := user.Current()
	die(err)
	segment.Value, err = os.Getwd()
	die(err)
	aliases = append(aliases, usr.HomeDir, icons["home"])

	for i := 0; i < len(aliases); i += 2 {
		prefix := filepath.Clean(aliases[i])
		icon := aliases[i+1]
		if strings.HasPrefix(segment.Value, prefix) {
			segment.Value = icon + segment.Value[len(prefix):]
			break
		}
	}

	length, _ := strconv.Atoi(os.Getenv("BRONZE_DIR_LENGTH"))
	if length != 0 {
		dirs := strings.Split(segment.Value, string(os.PathSeparator))
		for i := len(dirs) - 2; i >= 0; i-- {
			runes := []rune(dirs[i])
			if len(runes) > length {
				dirs[i] = string(runes[:length])
			}
		}

		segment.Value = strings.Join(dirs, string(os.PathSeparator))
	}
}
