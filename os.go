package main

import (
	"runtime"

	"github.com/go-ini/ini"
	. "github.com/reujab/bronze/types"
)

func osSegment(segment *Segment) {
	switch runtime.GOOS {
	case "darwin":
		segment.Value = icons["apple"]
	case "linux":
		var distro string
		file, err := ini.Load("/etc/os-release")
		if err == nil {
			distro = file.Section("").Key("ID").Value()
		}

		switch distro {
		case "arch":
			segment.Value = icons["arch"]
		case "centos":
			segment.Value = icons["centOS"]
		case "debian":
			segment.Value = icons["debian"]
		case "fedora":
			segment.Value = icons["fedora"]
		case "linuxmint":
			segment.Value = icons["mint"]
		case "suse", "opensuse":
			segment.Value = icons["SUSE"]
		case "ubuntu":
			segment.Value = icons["ubuntu"]
		case "elementary":
			segment.Value = icons["elementary"]
		default:
			segment.Value = icons["linux"]
		}
	case "freebsd", "netbsd", "openbsd":
		segment.Value = icons["bsd"]
	}
}
