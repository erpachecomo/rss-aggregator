package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login command needs one argument, the username")
	}
	user := cmd.Args[0]
	if err := s.config.SetUser(user); err != nil {
		return err
	}
	fmt.Printf("user '%v' has been set properly", user)
	return nil
}
