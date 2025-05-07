# Init
A very simple program to initialize a new project.


## Example config.toml
```toml
[[languages]]
name = "go"
directories = ["cmd", "cmd/main"]
files = ["cmd/main/main.go"]
shell_hook = [
    ["go", "mod", "init", "pwd"]
]
```

## TODO
- [x] Split the cli logic into its own file
- [x] Use pflag instead of if-else everywhere
- [x] Switch from JSON to TOML
- [x] Add shellhooks
- [ ] Remove the need for nested arrays in config.toml
- [ ] Add a way to split config files into multiple files
- [ ] Cleanup code
