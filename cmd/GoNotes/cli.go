//GoNotes
package main

import (
	"fmt"

	"github.com/mattackard/project-0/pkg/config"
	"github.com/mattackard/project-0/pkg/notes"
)

func main() {

	//determines what function to run based on the cli cmds
	switch config.Cmd {
	case "create":
		notes.CreateFile(config.Mycfg, config.FullPath, config.Open)
	case "config":
		notes.Config()
	case "list":
		myDir := notes.List(config.Mycfg.Paths.Notes)
		for _, v := range myDir {
			fmt.Println(v)
		}
	case "edit":
		notes.Edit(config.FullPath)
	case "delete":
		notes.Delete(config.FullPath)
	default:
		if config.Cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", config.Cmd)
		}
		notes.Print("help.txt")
	}
}