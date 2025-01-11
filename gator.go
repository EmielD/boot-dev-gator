package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/emield/gator/commands"
	"github.com/emield/gator/internal/config"
	"github.com/emield/gator/internal/database"
	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", config.Db_url)
	if err != nil {
		fmt.Printf("Error connectiong to database: %v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	stateVar := commands.State{Db: dbQueries, Config: &config}
	commandVar := commands.Command{Name: commandName, Arguments: os.Args[2:]}

	cmd := commands.Commands{CommandMap: make(map[string]func(*commands.State, commands.Command) error)}
	cmd.Register("login", commands.HandlerLogin)
	cmd.Register("register", commands.HandlerRegister)
	cmd.Register("reset", commands.HandlerReset)
	cmd.Register("users", commands.HandlerUsers)
	cmd.Register("agg", commands.HandlerAgg)
	cmd.Register("addfeed", commands.HandlerAddFeed)

	err = cmd.Run(&stateVar, commandVar)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
