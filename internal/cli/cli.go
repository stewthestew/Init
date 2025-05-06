package cli

import (
	"fmt"
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
	fmt.Println(cliOptions.List, cliOptions.Name)
	if cliOptions.Name == "" && cliOptions.List == false {
		flag.Usage()
		os.Exit(1) // Just incase since I don't know is flag.Usage() exits or not
	}

	return cliOptions
}
