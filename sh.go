package main

import (
	"fmt"
)

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
		return "%K{" + color + "}"
	case "bash":
		code, ok := colors["bg-"+color]
		if !ok {
			dief("invalid background color: %q", color)
		}
		return "\\[\x1b[" + code + "m\\]"
	default:
		code, ok := colors["bg-"+color]
		if !ok {
			dief("invalid background color: %q", color)
		}
		return "\x1b[" + code + "m"
	}
}

func escapeForeground(color string) string {
	switch shell {
	case "zsh":
		return "%F{" + color + "}"
	case "bash":
		code, ok := colors["fg-"+color]
		if !ok {
			dief("invalid foreground color: %q", color)
		}
		return "\\[\x1b[" + code + "m\\]"
	default:
		code, ok := colors["fg-"+color]
		if !ok {
			dief("invalid foreground color: %q", color)
		}
		return "\x1b[" + code + "m"
	}
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
