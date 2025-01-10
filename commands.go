package main

import (
	"fmt"

	"github.com/derjabineli/gator/internal/config"
	"github.com/derjabineli/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.cmds[cmd.name]
	if !exists{
		return fmt.Errorf("the provided command does not exist")
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}