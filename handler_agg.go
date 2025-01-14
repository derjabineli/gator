package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *state, cmd command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(rss)
	return nil
}
