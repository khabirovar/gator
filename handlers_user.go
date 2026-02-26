package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/khabirovar/gator/internal/config"
	"github.com/khabirovar/gator/internal/database"
	"strings"
	"time"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) <= 0 {
		return errors.New("Handler login expects a single argument, the username")
	}
	user, err := s.db.GetUser(context.Background(), strings.TrimSpace(cmd.args[0]))
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) <= 0 {
		return errors.New("Handler login expects a single argument, the username")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      strings.TrimSpace(cmd.args[0]),
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}
	return nil
}
