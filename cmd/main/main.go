package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"init/internal/cli"
	"init/internal/config"
)

func main() {
	home := os.Getenv("HOME")
	configFile := filepath.Join(home, ".config", "init", "config.json")
	configDir := filepath.Join(home, ".config", "init")
	created, err := config.CheckAndCreateConfig(configFile, configDir)

	if err != nil {
		log.Fatalf("Error creating config file %v", err)
	}

	if created {
		log.Println("Config file created")
	}

	opts := cli.ParseFlags()

	config, err := config.ParseConfig(configFile)
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
	for _, language := range config.Languages {
		if opts.Name == language.Name {
			for _, directory := range language.Directories {
				os.MkdirAll(directory, os.ModePerm)
			}
			for _, file := range language.Files {
				os.Create(file)
			}
		}

	}

}
