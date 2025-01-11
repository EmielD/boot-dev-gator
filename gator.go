package main

import (
	"database/sql"
	"fmt"
	"os"

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
	stateVar := state{db: dbQueries, config: &config}
	commandVar := command{name: commandName, arguments: os.Args[2:]}

	commands := commands{make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)

	err = commands.run(&stateVar, commandVar)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
