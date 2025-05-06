package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"init/internal/cli"
	iconfig "init/internal/config"
)

func main() {
	var index int
	if syscall.Getuid() == 0 {
		log.Fatalf("root is not supported")
	}

	home := os.Getenv("HOME")
	configFile := filepath.Join(home, ".config", "init", "config.toml")
	configDir := filepath.Join(home, ".config", "init")
	created, err := iconfig.CheckAndCreateConfig(configFile, configDir)

	if err != nil {
		log.Fatalf("Error creating config file %v", err)
	}

	if created {
		log.Println("Config file created")
	}

	opts := cli.ParseFlags()

	config, err := iconfig.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Error converting language index to int %v", err)
	}

	if opts.List {
		for i, language := range config.Languages {
			fmt.Printf("%v: Name: %v\n", i, language.Name)
		}
		os.Exit(0)
	}

	// Terrible code
	// but it works
	for i, language := range config.Languages {
		if opts.Name == language.Name {
			index = i
			for _, directory := range language.Directories {
				os.MkdirAll(directory, os.ModePerm)
			}
			for _, file := range language.Files {
				os.Create(file)
			}
		}

	}
	iconfig.RunShellHooks(config, index)

}
