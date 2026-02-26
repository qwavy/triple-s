package config

import (
	"errors"
	"flag"
)

type Config struct {
	Port int
	Dir  string
	Help bool
}

func UsageText() string {
	return `Simple Storage Service.

Usage:
  triple-s [-port <N>] [-dir <S>]
  triple-s --help

Options:
  --help     Show this screen.
  --port N   Port number
  --dir S    Path to the directory
`
}

func Parse(args []string) (Config, error) {
	port := flag.Int("port", 8080, "Port number")
	dir := flag.String("dir", "./data", "Path to the directory")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		return Config{Help: true}, nil
	}

	if *port < 1 || *port > 65535 {
		return Config{}, errors.New("wrong port number")
	}

	if *dir == "" {
		return Config{}, errors.New("directory is empty")
	}

	return Config{
		Port: *port,
		Dir:  *dir,
		Help: *help,
	}, nil
}
