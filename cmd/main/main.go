package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"init/internal/cli"
	"init/internal/config"
)

// TODO:
// Move cli args to its own file and use pflag
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

	if opts.Name == "" {
		fmt.Println("Please specify a language preset to initialize, or do --help for more information")
		os.Exit(1)
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

//
//	if err != nil {
//		log.Fatalf("Error parsing config file %v", err)
//	}
//
//	if err != nil {
//		log.Fatalf("Error converting language index to int %v", err)
//	}
//
//	//	if i > len(config.Languages) {
//	//		fmt.Printf("The language index %v does not exist.\n", i)
//	//		os.Exit(1)
//	//	}
//	//
//	//	for _, directory := range config.Languages[i].Directories {
//	//		os.MkdirAll(directory, os.ModePerm)
//	//	}
//
//	//	for _, file := range config.Languages[i].Files {
//	//		os.Create(file)
//	//	}
//
