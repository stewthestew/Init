package config

import (
	"encoding/json"
	"os"
)

const DefaultConfig = `
	{
    "languages": [
        {
            "name": "CHANGETHIS",
            "directories": ["test"],
            "files": ["test/test.txt"]
        }
    ]
	}
	`

type Config struct {
	DefaultConfigPath string
	DefaultConfigDir  string
	Languages         []Language `json:"languages"`
}

type Language struct {
	Name        string   `json:"name"`
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
}

func ParseConfig(fileName string) (Config, error) {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}
	var config Config

	err = json.Unmarshal([]byte(fileContents), &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func CheckAndCreateConfig(DefaultConfigPath string, DefaultConfigDir string) (bool, error) {
	_, err := os.ReadFile(DefaultConfigPath)
	if err != nil {
		os.Mkdir(DefaultConfigDir, os.ModePerm)
		os.Create(DefaultConfigPath)
		file, err := os.OpenFile(DefaultConfigPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return false, err
		}
		file.WriteString(DefaultConfig)
		return true, nil
	}
	return false, nil
}
