package main

import (
	"github.com/khabirovar/gator/internal/config"
	"github.com/khabirovar/gator/internal/database"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	st := state{
		db: database.New(db),
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	args := os.Args 
	if len(args) <= 2 {
		log.Fatal("command have no arguments")
	}
	
	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
