package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var hexRegex = regexp.MustCompile(`^[a-f\d]{6}$`)

var colors = map[string]string{
	"bg-none":    "0",
	"fg-black":   "30",
	"fg-red":     "31",
	"fg-green":   "32",
	"fg-yellow":  "33",
	"fg-blue":    "34",
	"fg-magenta": "35",
	"fg-cyan":    "36",
	"fg-white":   "37",
	"bg-black":   "40",
	"bg-red":     "41",
	"bg-green":   "42",
	"bg-yellow":  "43",
	"bg-blue":    "44",
	"bg-magenta": "45",
	"bg-cyan":    "46",
	"bg-white":   "47",
}

func escapeBackground(color string) string {
	switch shell {
	case "zsh":
		// 24-bit color
		if hexRegex.MatchString(color) {
			return "%{\x1b[48;2;" + escapeHex(color) + "m%}"
		}

		// 16 and 256 colors
		return "%K{" + color + "}"
	case "bash":
		// 16 colors
		code := colors["bg-"+color]
		if code != "" {
			return "\\[\x1b[" + code + "m\\]"
		}

		// 256 colors
		_, err := strconv.Atoi(color)
		if err == nil {
			return "\\[\x1b[48;5;" + color + "m\\]"
		}

		// 24-bit color
		if hexRegex.MatchString(color) {
			return "\\[\x1b[48;2;" + escapeHex(color) + "m\\]"
		}

		dief("invalid background color: %q", color)
	default:
		// 16 colors
		code := colors["bg-"+color]
		if code != "" {
			return "\x1b[" + code + "m"
		}

		// 24-bit color
		if hexRegex.MatchString(color) {
			return "\x1b[48;2;" + escapeHex(color) + "m"
		}

		// 256 colors
		_, err := strconv.Atoi(color)
		if err == nil {
			return "\x1b[48;5;" + color + "m"
		}

		dief("invalid background color: %q", color)
	}

	return ""
}

func escapeForeground(color string) string {
	switch shell {
	case "zsh":
		// 24-bit color
		if hexRegex.MatchString(color) {
			return "%{\x1b[38;2;" + escapeHex(color) + "m%}"
		}

		// 16 and 256 colors
		return "%F{" + color + "}"
	case "bash":
		// 16 colors
		code := colors["fg-"+color]
		if code != "" {
			return "\\[\x1b[" + code + "m\\]"
		}

		// 256 colors
		_, err := strconv.Atoi(color)
		if err == nil {
			return "\x1b[38;5;" + color + "m"
		}

		// 24-bit color
		if hexRegex.MatchString(color) {
			return "\\[\x1b[38;2;" + escapeHex(color) + "m\\]"
		}

		dief("invalid foreground color: %q", color)
	default:
		// 16 colors
		code := colors["fg-"+color]
		if code != "" {
			return "\x1b[" + code + "m"
		}

		// 256 colors
		_, err := strconv.Atoi(color)
		if err == nil {
			return "\x1b[38;5;" + color + "m"
		}

		// 24-bit color
		if hexRegex.MatchString(color) {
			return "\x1b[38;2;" + escapeHex(color) + "m"
		}

		dief("invalid foreground color: %q", color)
	}

	return ""
}

func resetColors() {
	switch shell {
	case "zsh":
		fmt.Print("%{%f%}")
	case "bash":
		fmt.Print("\\[\x1b[0m\\]")
	default:
		fmt.Print("\x1b[0m")
	}
}

func escapeHex(color string) string {
	var r, g, b byte
	fmt.Sscanf(color, "%02x%02x%02x", &r, &g, &b)
	return fmt.Sprintf("%d;%d;%d", r, g, b)
}
