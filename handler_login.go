package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login command needs one argument, the username")
	}
	user := cmd.Args[0]
	_, err := s.db.GetUserByName(context.Background(), user)

	if err != nil {
		log.Fatalf("user '%v' does not exist in the database", user)
	}

	if err := s.config.SetUser(user); err != nil {
		return err
	}
	fmt.Printf("user '%v' has been set properly", user)
	return nil
}
