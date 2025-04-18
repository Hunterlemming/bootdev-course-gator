package main

import (
	"fmt"
	"os"

	"github.com/hunterlemming/bootdev-course-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	s := state{cfg: &cfg}

	comms := commands{handlers: make(map[string]func(*state, command) error)}
	comms.register("login", handlerLogin)

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
