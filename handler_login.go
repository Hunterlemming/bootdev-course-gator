package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("username argument required")
	}

	userName := cmd.args[0]
	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user: %s", err.Error())
	}

	fmt.Printf("%s logged in\n", s.cfg.CurrentUserName)
	return nil
}
