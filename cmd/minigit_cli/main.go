package main

import (
	"fmt"
	"minigit/cmd/minigit_cli/internal/repo"
	"minigit/cmd/minigit_cli/utils"
	"os"
)

const (
	INIT = "init"
	ADD = "add"
)

func main(){
	if len(os.Args) <= 1{
		fmt.Println("start a working area")
		fmt.Println("\tinit \tCreate an empty Mini Git repository or reinitialize an existing one")
	}
	args := utils.New(os.Args)
	cmd, _ := args.Next()
	switch cmd {
		case INIT:
			repo.Init()
		case ADD:
			repo.Add()
		default:
			printNotRecongizedCommandError(cmd)
			
	}
}

func printNotRecongizedCommandError(cmd string){
	fmt.Printf("minigit: '%s' is not a command. See 'minigit --help' for more info.\n", cmd)
}
