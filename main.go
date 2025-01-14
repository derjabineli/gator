package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/derjabineli/gator/internal/config"
	"github.com/derjabineli/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config. error: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("cant open database connection. error: %v", err)
	}

	dbQueries := database.New(db)

	programState := state{cfg: &cfg, db: dbQueries}
	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAggregate)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerGetFeeds)
	

	args := os.Args

	if len(args) == 1 {
		fmt.Println("error: not enough arguments provided") 
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args}

	err = commands.run(&programState, cmd)
	if err != nil {
		fmt.Printf("\033[0;31m error:\033[0m %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}