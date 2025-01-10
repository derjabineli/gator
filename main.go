package main

import (
	"fmt"
	"os"

	"github.com/derjabineli/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config. error: %v\n", err)
	}

	state := State{config: &cfg}
	commands := Commands{
		cmds: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) == 1 {
		fmt.Println("error: not enough arguments provided") 
		os.Exit(1)
	}

	cmd := Command{name: args[1], args: args}

	err = commands.run(&state, cmd)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}