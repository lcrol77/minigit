package main

import (
	"fmt"
	"minigit/cmd/minigit_cli/internal/repo"
	"os"
)

const (
	INIT = "init"
)

func main(){
	if len(os.Args) <= 1{
		fmt.Println("start a working area")
		fmt.Println("\tinit \tCreate an empty Mini Git repository or reinitialize an existing one")
	}
	cmd := os.Args[1]
	switch cmd {
		case INIT:
			repo.Init()
	}
}
