package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"init/internal/config"
)

// TODO:
// Move cli args to its own file and use pflag
func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Usage: %s <Language index>", args[0])
	}


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

	if args[1] == "list" {
		fmt.Println("Available languages:")
		for i, language := range config.Languages {
			fmt.Printf("%v: %v\n", i, language.Name)
		}
		os.Exit(0)
	}
	if err != nil {
		log.Fatalf("Error parsing config file %v", err)
	}
	i, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatalf("Error converting language index to int %v", err)
	}

	if i > len(config.Languages) {
		fmt.Printf("The language index %v does not exist.\n", i, )
		os.Exit(1)
	}

	for _, directory := range config.Languages[i].Directories {
		os.MkdirAll(directory, os.ModePerm)
	}

	for _, file := range config.Languages[i].Files {
		os.Create(file)
	}
}
