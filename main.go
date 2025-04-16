package main

import (
	"fmt"

	"github.com/hunterlemming/bootdev-course-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg.SetUser("Krisz")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cfg)
}
