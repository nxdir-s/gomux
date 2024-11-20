# GoMux

GoMux is a CLI tool that helps with setting up tmux for a project

## Installation

Install GoMux with the following

```shell
$ go install github.com/nxdir-s/gomux/cmd/cli@latest
```

## Usage

GoMux is intended to be used to setup your tmux windows for a project. While in your project directory
run `gomux` to automatically setup tmux using a config file

### Configuration

GoMux can be configured using a toml file named `.gomux.toml`

#### Example

```toml
title = "Example GoMux Config"

[config]
session = "SessionName"
project = "~/projects/myproject"
start_index = 0

[windows]

[windows.editor]
name = "editor"
cmd = "vim ."

[windows.docker]
name = "docker"
cmd = "docker compose up"

[windows.database]
name = "database"
cmd = "psql"

[windows.server]
name = "server"
cmd = "go run cmd/main.go"
```
