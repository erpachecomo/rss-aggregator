package main

import (
	"database/sql"
	"log"
	"os"
	config "rss-aggregator/internal"
	"rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		log.Fatalf("Error reaching the db: %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		config: &cfg,
		db:     dbQueries,
	}

	commands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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
