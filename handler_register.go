package main

import (
	"context"
	"errors"
	"log"
	"rss-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register command needs name arg")
	}
	name := cmd.Args[0]
	_, err := s.db.GetUserByName(context.Background(), name)
	if err == nil {
		log.Fatal("User already exists")
	}
	params := database.CreateUserParams{
		Name:      name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)

	}
	if err = s.config.SetUser(user.Name); err != nil {
		return err
	}
	log.Printf("User created successfully: %v", user)
	return nil
}
