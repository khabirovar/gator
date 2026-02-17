package main

import (
	"errors"
	"fmt"
	"time"
	"context"
	"strings"
  "github.com/google/uuid"
	"github.com/khabirovar/gator/internal/config"
	"github.com/khabirovar/gator/internal/database"
)

type state struct {
	db *database.Queries
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
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: strings.TrimSpace(cmd.args[0]),
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
