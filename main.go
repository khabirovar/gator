package main

import (
	"github.com/khabirovar/gator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	st := state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args 
	if len(args) <= 2 {
		log.Fatal("command have no arguments")
	}
	
	cmd := command{
		name: args[1],
		args: args[2:],
	}

	cmds.run(&st, cmd)
}
