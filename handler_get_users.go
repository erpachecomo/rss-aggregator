package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())

	if err != nil {
		return err
	}

	for _, v := range users {

		if v.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", v.Name)
		} else {
			fmt.Printf("* %v\n", v.Name)
		}
	}
	return nil
}
