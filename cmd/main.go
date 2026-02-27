package main

import (
	"fmt"
	"os"

	config2 "triple-s/internal/config"
	"triple-s/internal/server"
)

func main() {
	cfg, err := config2.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	if cfg.Help {
		fmt.Println(config2.UsageText())
		return
	}

	err = server.RunServer(cfg)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
