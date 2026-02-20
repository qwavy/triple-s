package config

import "flag"

type Config struct {
	Port string
	Dir  string
}

func Load() *Config {
	port := flag.String("port", "./data", "Port number")
	dir := flag.String("dir", "./data", "Path to the directory")

	var cfg Config

}
