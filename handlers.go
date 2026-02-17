package main

import (
	"errors"
	"fmt"
	"github.com/khabirovar/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) <= 0 {
		return errors.New("Handler login expects a single argument, the username") 
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
