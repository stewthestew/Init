package main

import (
	"log"
	"os"
	"path/filepath"

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

	config, err := config.ParseConfig(configFile)

	if err != nil {
		log.Fatalf("Error parsing config file %v", err)
	}
	for _, language := range config.Languages {
		for _, directory := range language.Directories {
			os.MkdirAll(directory, os.ModePerm)
		}

		for _, file := range language.Files {
			os.Create(file)
		}
	}
}
