package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/derjabineli/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	savePosts(s, *rss, feed.ID)
	return nil
}

func savePosts(s *state, rss RSSFeed, feedId uuid.UUID) {
	ctx := context.Background()
	for _, item := range rss.Channel.Item {
		publishDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		var dbPublishDate sql.NullTime

		if err != nil {
			dbPublishDate = sql.NullTime{Valid: false}
		} else {
			dbPublishDate = sql.NullTime{Time: publishDate, Valid: true}
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Title: sql.NullString{String: item.Title, Valid: true}, Url: item.Link, Description: sql.NullString{String: item.Description, Valid: true}, PublishedAt: dbPublishDate, FeedID: feedId})

		var pgErr *pq.Error
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				continue
			} else {
				fmt.Printf("Encountered Error: %v\n", err)
			}
		}
	}
}