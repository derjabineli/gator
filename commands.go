package main

import (
	"fmt"

	"github.com/derjabineli/gator/internal/config"
)

type State struct {
	config *config.Config
}

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmds map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmds[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
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

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("you must provide a username")
	}

	userName := cmd.args[2]

	err := s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("The User has successfully been set to %s\n", userName)
	return nil
}