package main

import (
	"fmt"
	"os"
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
	case "fish":
		code, ok := colors["bg-"+color]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid background color: %q.", color)
			os.Exit(1)
		}
		// fish is smart enough to detect zero-width characters
		return "\x1b[" + code + "m"
	case "zsh":
		return "%K{" + color + "}"
	default:
		code, ok := colors["bg-"+color]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid background color: %q.", color)
			os.Exit(1)
		}
		// the brackets tell bash the escape code is zero-width
		return "\\[\x1b[" + code + "m\\]"
	}
}

func escapeForeground(color string) string {
	switch shell {
	case "fish":
		code, ok := colors["fg-"+color]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid foreground color: %q.", color)
			os.Exit(1)
		}
		return "\x1b[" + code + "m"
	case "zsh":
		return "%F{" + color + "}"
	default:
		code, ok := colors["fg-"+color]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid foreground color: %q.", color)
			os.Exit(1)
		}
		return "\\[\x1b[" + code + "m\\]"
	}
}

func resetColors() {
	switch shell {
	case "fish":
		fmt.Printf("\x1b[0m")
	case "zsh":
		fmt.Print("%{%f%}")
	default:
		fmt.Printf("\\[\x1b[0m\\]")
	}
}
