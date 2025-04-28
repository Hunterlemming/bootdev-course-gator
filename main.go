package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/hunterlemming/bootdev-course-gator/internal/config"
	"github.com/hunterlemming/bootdev-course-gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	s := state{
		db:  database.New(db),
		cfg: &cfg,
	}

	comms := commands{handlers: make(map[string]func(*state, command) error)}
	comms.register("login", handlerLogin)
	comms.register("register", handlerRegister)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = comms.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
