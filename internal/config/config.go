package config

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

const DefaultConfig = `
[[languages]]
name = "CHANGETHIS"
directories = ["CHANGETHAT_IN_CONFIG"]
files = ["CHANGETHIS_IN_CONFIG.txt"]
`

type Config struct {
	DefaultConfigPath string
	DefaultConfigDir  string
	Languages         []Language `toml:"languages"`
}

type Language struct {
	Name        string     `toml:"name"`
	Directories []string   `toml:"directories"`
	Files       []string   `toml:"files"`
	ShellHooks  [][]string `toml:"shell_hook"`
}

func ParseConfig(fileName string) (Config, error) {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}
	var config Config

	err = toml.Unmarshal([]byte(fileContents), &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func RunShellHooks(config Config, index int) error {
	language := config.Languages[index]

	for _, hook := range language.ShellHooks {
		if len(hook) == 0 {
			fmt.Println("No shell hook specified. Skipping...")
			return nil
		}

		cmd := exec.Command(hook[0], hook[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
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
