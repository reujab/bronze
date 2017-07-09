package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-ini/ini"
)

func osSegment(segment *segment) {
	if runtime.GOOS == "darwin" {
		segment.value = iconApple
	} else {
		file, err := ini.Load("/etc/os-release")
		check(err)
		switch file.Section("").Key("ID").Value() {
		case "arch":
			segment.value = iconArch
		case "centos":
			segment.value = iconCentOS
		case "debian":
			segment.value = iconDebian
		case "fedora":
			segment.value = iconFedora
		case "linuxmint":
			segment.value = iconMint
		case "suse", "opensuse":
			segment.value = iconSUSE
		case "ubuntu":
			segment.value = iconUbuntu
		case "elementary":
			segment.value = iconElementary
		default:
			fmt.Fprintln(os.Stderr, "bronze: Unrecognized operating system.")
		}
	}
	segment.visible = true
}
