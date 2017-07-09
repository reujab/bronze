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
			// TODO
		case "zsh":
			fmt.Println(`unsetopt prompt_subst; precmd() { PROMPT="$(bronze print "${BRONZE[@]}") "; }`)
		case "bash":
			fmt.Println(`PS1='$(bronze print "${BRONZE[@]}") '`)
		default:
			fmt.Print(`echo "bronze: Unrecognized shell."`)
		}
		return nil
	},
}
