package main

import (
	"log"
	"os"
	config "rss-aggregator/internal"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}

	programState := &state{
		config: &cfg,
	}

	commands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	a := os.Args

	if len(a) < 2 {
		log.Fatal("insufficient arguments")
	}
	cmd := command{
		Name: "login",
		Args: a[2:],
	}
	if err := commands.run(programState, cmd); err != nil {
		log.Fatal(err)
	}
}
