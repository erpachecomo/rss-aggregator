package main

import (
	"fmt"
	"log"
	config "rss-aggregator/internal"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}

	err = cfg.SetUser("erpachecomo")
	if err != nil {
		log.Printf("Error setting user in the config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}
	fmt.Println(cfg)

}
