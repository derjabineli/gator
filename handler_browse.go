package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/derjabineli/gator/internal/database"
)

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 2 {
		converted, err := strconv.Atoi(cmd.args[2])
		if err == nil {
			limit = converted
		}
	}

	ctx := context.Background()
	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("could not find posts. details: %v", err)
	}

	for _, post := range posts {
		line := strings.Repeat("-", len(post.Title.String))
		fmt.Println(Cyan + post.Title.String + Reset)
		fmt.Println(line)
		fmt.Println(Green + post.Url + Reset)
		fmt.Println(post.Description.String)
	}
	return nil
}