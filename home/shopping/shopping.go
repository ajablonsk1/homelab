package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajablonsk1/homelab/home/shopping/internal/command"
	"github.com/ajablonsk1/homelab/home/shopping/internal/parse"
)

func main() {
	var err error
	args := os.Args

	if len(args) > 1 {
		err = executeCommand(args)
	} else {
		err = fmt.Errorf("no command entered! see --help for more info")
	}

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func executeCommand(args []string) error {
	cmd := args[1]
	switch cmd {
	case "ls":
		return command.Ls()
	case "add":
		product, err := parse.AddProductArguments()
		if err != nil {
			return err
		}
		return command.Add(&product)
	case "update":
		name := args[2]
		product, err := parse.UpdateProductArguments()
		if err != nil {
			return err
		}
		return command.Update(name, &product)
	case "rm":
		name := args[2]
		return command.Rm(name)
	default:
		return fmt.Errorf("unrecognized command! Usage shopping <command> <product>")
	}
}
