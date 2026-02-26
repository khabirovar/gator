package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("We don't have function '%s' in register", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.handlers[name]; ok {
		return fmt.Errorf("You try overwrite command '%s' in command register", name)
	}
	c.handlers[name] = f

	return nil
}
