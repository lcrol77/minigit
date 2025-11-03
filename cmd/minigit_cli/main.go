package main

import (
	"fmt"
	"minigit/cmd/minigit_cli/internal"
	"minigit/cmd/minigit_cli/utils/args"
	"os"
)

const (
	INIT     = "init"
	ADD      = "add"
	CAT_FILE = "cat-file"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("start a working area")
		fmt.Println("\tinit \tCreate an empty Mini Git repository or reinitialize an existing one")
	}
	args := args.New(os.Args)
	cmd, _ := args.Next()
	switch cmd {
	case INIT:
		repo.Init()
	case ADD:
		for args.HasNext() {
			cmd, _ := args.Next()
			repo.Add(cmd)
		}
	default:
		printNotRecongizedCommandError(cmd)
	}
}

func printNotRecongizedCommandError(cmd string) {
	fmt.Printf("minigit: '%s' is not a command. See 'minigit --help' for more info.\n", cmd)
}
