package main

import (
	"fmt"
	"os"

	"github.com/emield/gator/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter a command")
		os.Exit(1)
	}

	commandName := os.Args[1]

	config, err := config.Read()
	if err != nil {
		fmt.Printf("Something went wrong reading the config: %v\n", err)
		os.Exit(1)
	}

	stateVar := state{config: &config}
	commandVar := command{name: commandName, arguments: os.Args[2:]}

	commands := commands{make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)

	err = commands.run(&stateVar, commandVar)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
