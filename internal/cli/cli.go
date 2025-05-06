package cli

import (
	"os"

	flag "github.com/spf13/pflag"
)

type CliOptions struct {
	List bool
	Name string
}

func ParseFlags() CliOptions {
	var cliOptions CliOptions
	flag.BoolVarP(&cliOptions.List, "list", "l", false, "List available languages")
	flag.StringVarP(&cliOptions.Name, "name", "n", "", "Name of the language preset to initialize")
	flag.Parse()
	if cliOptions.Name == "" {
		flag.Usage()
		os.Exit(1) // Just incase since I don't know is flag.Usage() exits or not
	}

	return cliOptions
}
