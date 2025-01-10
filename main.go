package main

import (
	"fmt"

	"github.com/derjabineli/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config. error: %v\n", err)
	}
	
	err = cfg.SetUser("eli")
	if err != nil {
		fmt.Printf("error setting user. error: %v\n", err)
	}
	
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config. error: %v\n", err)
	}
}