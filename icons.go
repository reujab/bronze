package main

import (
	"os"
	"strings"
)

var separator = os.Getenv("BRONZE_SEPARATOR")
var thinSeparator = os.Getenv("BRONZE_THIN_SEPARATOR")
var iconMode = os.Getenv("BRONZE_ICONS")
var icons = map[string]string{}

func init() {
	// set defaults
	if separator == "" {
		separator = "\ue0b0" // Powerline
	}
	if thinSeparator == "" {
		thinSeparator = "\ue0b1" // Powerline
	}
	if iconMode == "" {
		iconMode = "nerd"
	}

	switch iconMode {
	case "nerd":
		initIcon("apple", "\uf179")      // Font Awesome; apple
		initIcon("arch", "\uf300")       // Font Linux
		initIcon("centOS", "\uf301")     // Font Linux
		initIcon("debian", "\uf302")     // Font Linux
		initIcon("fedora", "\uf303")     // Font Linux
		initIcon("mint", "\uf304")       // Font Linux
		initIcon("SUSE", "\uf308")       // Font Linux
		initIcon("ubuntu", "\uf30c")     // Font Linux
		initIcon("elementary", "\uf311") // Font Linux
		initIcon("linux", "\uf17c")      // Font Awesome; linux
		initIcon("bsd", "\uf30e")        // Font Linux
		initIcon("root", "\ue00a")       // Pomicons; lightning
		initIcon("readonly", "\uf023")   // Font Awesome; lock
		initIcon("failed", "\ue009")     // Pomicons; exclamation
		initIcon("job", "\ue615")        // Seti UI
		initIcon("package", "\uf187")    // Font Awesome; archive
		initIcon("home", "\uf015")       // Font Awesome; home
		initIcon("github", "\uf09b")     // Font Awesome; github
		initIcon("gitlab", "\uf296")     // Font Awesome; gitlab
		initIcon("bitbucket", "\uf171")  // Font Awesome; bitbucket
		initIcon("git", "\ue0a0")        // Powerline
		initIcon("stash", "\uf01c")      // Font Awesome; inbox
		initIcon("ahead", "\uf148")      // Font Awesome; level-up
		initIcon("behind", "\uf149")     // Font Awesome; level-down
		initIcon("modified", "\uf111")   // Unicode
		initIcon("staged", "\uf067")     // Unicode
	case "unicode":
		// TODO: test if it's possible to use \uf8ff on an Apple machine
		initIcon("apple", "\U0001f34e") // Emoji; red apple
		initIcon("arch", "")
		initIcon("centOS", "")
		initIcon("debian", "")
		initIcon("fedora", "")
		initIcon("mint", "")
		initIcon("SUSE", "")
		initIcon("ubuntu", "")
		initIcon("elementary", "")
		initIcon("linux", "\U0001f427")    // Emoji; penguin
		initIcon("bsd", "\U0001f608")      // Emoji; smiling face with horns
		initIcon("root", "\u26a1")         // Emoji; high voltage
		initIcon("readonly", "\U0001f512") // Emoji; locked
		initIcon("failed", "\u2757")       // Emoji; exclamation mark
		initIcon("job", "\u2699")          // Emoji; gear
		initIcon("package", "\U0001f4e6")  // Emoji; package
		initIcon("home", "\U0001f3e0")     // Emoji; house
		initIcon("github", "")
		initIcon("gitlab", "")
		initIcon("bitbucket", "")
		initIcon("git", "")
		initIcon("stash", "\U0001f4e5") // Emoji; inbox tray
		initIcon("ahead", "\u2191")     // Unicode
		initIcon("behind", "\u2193")    // Unicode
		initIcon("modified", "\u25cf")  // Unicode
		initIcon("staged", "\u271a")    // Unicode
	case "ascii":
		initIcon("apple", "")
		initIcon("arch", "")
		initIcon("centOS", "")
		initIcon("debian", "")
		initIcon("fedora", "")
		initIcon("mint", "")
		initIcon("SUSE", "")
		initIcon("ubuntu", "")
		initIcon("elementary", "")
		initIcon("linux", "")
		initIcon("bsd", "")
		initIcon("root", "#")
		initIcon("readonly", "@")
		initIcon("failed", "!")
		initIcon("job", "&")
		initIcon("package", "")
		initIcon("home", "~")
		initIcon("github", "")
		initIcon("gitlab", "")
		initIcon("bitbucket", "")
		initIcon("git", "")
		initIcon("stash", "#")
		initIcon("ahead", ">")
		initIcon("behind", "<")
		initIcon("modified", "*")
		initIcon("staged", "+")
	}
}

func initIcon(name, defaultValue string) {
	env := os.Getenv("BRONZE_ICON_" + strings.ToUpper(name))
	if env == "" {
		icons[name] = defaultValue
	} else {
		icons[name] = env
	}
}
