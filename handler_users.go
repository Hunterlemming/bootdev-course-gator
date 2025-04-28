package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hunterlemming/bootdev-course-gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("username argument required")
	}

	userName := cmd.args[0]
	ctx := context.Background()
	createParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	createdUser, err := s.db.CreateUser(ctx, createParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User '%s' was registered and set.", userName)
	printUser(createdUser)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("username argument required")
	}

	userName := cmd.args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("%s logged in\n", s.cfg.CurrentUserName)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
