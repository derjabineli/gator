package main

import (
	"context"
	"fmt"
	"time"

	"github.com/derjabineli/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 4 {
		return fmt.Errorf("you must provide a rss feed name and url") 
	}

	ctx := context.Background()

	name := cmd.args[2]
	url := cmd.args[3]
	
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: user.ID})
	if err != nil {
		return fmt.Errorf("create feed failed %v", err)
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("could not subscribe to feed")
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
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("could not retrieve rss feeds")
		}
		fmt.Println(feed.Name)
		fmt.Printf("  -URL: %v\n", feed.Url)
		fmt.Printf("  -Created By: %v\n", user.Name)
	}
	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("must provide a rss feed url")
	}
	url := cmd.args[2]
	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("feed does not exist")
	}

	followFeed, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("could not subscribe to feed")
	}

	fmt.Printf("%v subscribed to %v\n", followFeed.UserName, followFeed.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("could not find followed feeds")
	}

	fmt.Println("Followed Feeds:")
	fmt.Println("---------------------")
	for _, feed := range feeds {
		fmt.Printf("  - %v\n", feed.Name.String)
	}
	return nil
}