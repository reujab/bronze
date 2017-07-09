package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdInit = cli.Command{
	Name:        "init",
	Usage:       "Initializes the shell for use of bronze",
	Description: `Put 'eval "$(bronze init)"' in your ~/.*shrc`,
	Action: func(ctx *cli.Context) error {
		switch shell {
		case "fish":
			fmt.Println(`function fish_prompt; env CMDTIME={$CMD_DURATION}ms bronze print $BRONZE; echo -n ' '; end`)
		case "zsh":
			fmt.Println(`BRONZE_START=$(date +%s%3N)
unsetopt prompt_subst

preexec() {
	BRONZE_START=$(date +%s%3N)
}

precmd() {
	PROMPT="$(CMDTIME=$(($(date +%s%3N)-$BRONZE_START))ms bronze print "${BRONZE[@]}") "
}`)
		case "bash":
			fmt.Println(`PS1='$(bronze print "${BRONZE[@]}") '`)
		default:
			fmt.Print(`echo "bronze: Unrecognized shell."`)
		}
		return nil
	},
}
