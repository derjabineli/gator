package main

import (
	"context"
	"fmt"
	"time"

	"github.com/derjabineli/gator/internal/database"
	"github.com/google/uuid"
)
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("you must provide a username")
	}

	userName := cmd.args[2]

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("username not found")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("The User has successfully been set to %s\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("you must provide a username")
	}

	userName := cmd.args[2]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: userName})
	if err != nil {
		return fmt.Errorf("username not available")
	}

	s.cfg.SetUser(user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}