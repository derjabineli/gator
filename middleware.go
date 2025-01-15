package main

import (
	"context"
	"fmt"

	"github.com/derjabineli/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
        if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("please log in or register before continuing")
		}

		ctx := context.Background()
		user, err := s.db.GetUserByName(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("please log in or register before continuing")
		}

		err = handler(s, cmd, user)
		if err != nil {
			return fmt.Errorf("please log in or register before continuing")
		}
		return nil
    }
}