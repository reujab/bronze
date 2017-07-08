package main

func escapeBackground(color string) string {
	switch shell {
	case "zsh":
		return "%K{" + color + "}"
	default:
		// TODO: default to bash
		return ""
	}
}

func escapeForeground(color string) string {
	switch shell {
	case "zsh":
		return "%F{" + color + "}"
	default:
		// TODO: default to bash
		return ""
	}
}
