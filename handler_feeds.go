package main

import (
	"context"
	"fmt"
	"time"

	"github.com/derjabineli/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 4 {
		return fmt.Errorf("you must provide a rss feed name and url") 
	}
	
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}

	name := cmd.args[2]
	url := cmd.args[3]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID})
	if err != nil {
		return fmt.Errorf("create feed failed %v", err)
	}

	fmt.Print(feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve rss feeds")
	}

	fmt.Println("Feeds:")
	fmt.Println("------------------------------")
	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Printf("  -URL: %v\n", feed.Url)
		fmt.Printf("  -Created By: %v\n", feed.UserName.String)
	}
	return nil
}