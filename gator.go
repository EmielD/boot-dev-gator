package main

import (
	"fmt"

	"github.com/emield/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Something went wrong reading the config: %v\n", err)
	}

	config, err = config.SetUser("Emiel")
	if err != nil {
		fmt.Printf("Something went wrong setting the user: %v\n", err)
	}

	fmt.Printf("%s\n", config)
}
