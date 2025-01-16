package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/derjabineli/gator/internal/database"
)

func handlerAggregate(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 3 {
		return fmt.Errorf("please provide a time duration that you'd like to wait between requests")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[2])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s, user)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state, user database.User) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed. details: %v", err)
	}

	currentTime := sql.NullTime{Time: time.Now(), Valid: true}
	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{ID: feed.ID, LastFetchedAt: currentTime, UpdatedAt: time.Now()})
	if err != nil {
		return fmt.Errorf("something went wrong. details: %v", err)
	}

	rss, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch rss feed. details: %v", err)
	}

	printRSSItems(*rss)
	return nil
}

func printRSSItems(rss RSSFeed) {
	fmt.Printf("%v\n", rss.Channel.Title)

	for _, item := range rss.Channel.Item {
		fmt.Println(item.Title)
	}

	fmt.Println()
}