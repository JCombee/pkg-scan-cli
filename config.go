package main

import (
	"os"
)

type Config struct {
	FileRoot string
}

func ConfigDefault() Config {
	return Config{
		FileRoot: ".",
	}
}

func (c *Config) InitArgs() {
	args := os.Args
	if len(args) > 1 {
		c.FileRoot = args[1]
	}
}
