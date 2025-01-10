package main

import (
	"fmt"
	"testing"

	"github.com/derjabineli/gator/internal/config"
)

func TestCommandInterface(t *testing.T) {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config. error: %v\n", err)
	}

	state := State{config: &cfg}
	commands := Commands{
		cmds: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handlerLogin)

	cases := []struct {
		name     string
		input    []string
		expected error
	}{
		{
			name:     "Missing username",
			input:    []string{"", "login"},
			expected: fmt.Errorf("you must provide a username"),
		},
		{
			name:     "Valid username",
			input:    []string{"", "login", "eli"},
			expected: nil,
		},
		{
			name:     "Invalid command",
			input:    []string{"", "register", "eli"},
			expected: fmt.Errorf("the provided command does not exist"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cmd := Command{name: c.input[1], args: c.input}
			err := commands.run(&state, cmd)

			if !errorsAreEqual(err, c.expected) {
				t.Errorf("Test case %q failed: input %v, expected error: %v, got: %v", c.name, c.input, c.expected, err)
			}
		})
	}
}

func errorsAreEqual(err1, err2 error) bool {
	if err1 == nil || err2 == nil {
		return err1 == err2
	}
	return err1.Error() == err2.Error()
}