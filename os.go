package main

import (
	"runtime"

	"github.com/go-ini/ini"
)

func osSegment(segment *segment) {
	switch runtime.GOOS {
	case "darwin":
		segment.value = icons["apple"]
	case "linux":
		file, err := ini.Load("/etc/os-release")
		check(err)
		switch file.Section("").Key("ID").Value() {
		case "arch":
			segment.value = icons["arch"]
		case "centos":
			segment.value = icons["centOS"]
		case "debian":
			segment.value = icons["debian"]
		case "fedora":
			segment.value = icons["fedora"]
		case "linuxmint":
			segment.value = icons["mint"]
		case "suse", "opensuse":
			segment.value = icons["SUSE"]
		case "ubuntu":
			segment.value = icons["ubuntu"]
		case "elementary":
			segment.value = icons["elementary"]
		}
		if segment.value == "" {
			segment.value = icons["linux"]
		}
	}
}
