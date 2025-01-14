package main

import (
	"database/sql"
	"fmt"
	"os"

	c "github.com/emield/gator/commands"
	"github.com/emield/gator/internal/config"
	"github.com/emield/gator/internal/database"
	m "github.com/emield/gator/internal/middleware"
	"github.com/emield/gator/internal/types"
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
	stateVar := types.State{Db: dbQueries, Config: &config}
	commandVar := types.Command{Name: commandName, Arguments: os.Args[2:]}

	cmd := c.Commands{CommandMap: make(map[string]func(*types.State, types.Command) error)}
	cmd.Register("login", c.HandlerLogin)
	cmd.Register("register", c.HandlerRegister)
	cmd.Register("reset", c.HandlerReset)
	cmd.Register("users", c.HandlerUsers)
	cmd.Register("agg", c.HandlerAgg)
	cmd.Register("addfeed", m.LoggedIn(c.HandlerAddFeed))
	cmd.Register("feeds", c.HandlerFeeds)
	cmd.Register("follow", m.LoggedIn(c.HandlerFollow))
	cmd.Register("following", m.LoggedIn(c.HandlerFollowing))
	cmd.Register("unfollow", m.LoggedIn(c.HandleUnfollow))
	cmd.Register("browse", m.LoggedIn(c.HandlerBrowse))

	err = cmd.Run(&stateVar, commandVar)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
