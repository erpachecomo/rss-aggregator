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

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	if err := commands.run(programState, cmd); err != nil {
		log.Fatal(err)
	}
}
