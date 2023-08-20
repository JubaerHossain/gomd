package gomd

import (
	"github.com/JubaerHossain/gomd/command"
	"github.com/JubaerHossain/gomd/config"
	"github.com/spf13/cobra"
)

var asciiArt = `
     ____ _____  ____ ___  ____/ /
    / __ / __ \/ __  __ \/ __  / 
   / /_/ / /_/ / / / / / / /_/ /  
   \__, /\____/_/ /_/ /_/\__,_/   
  /____/
  
  A CLI tool to building restful API with Go
`

// Create a Cobra command for gomd
var cmd = &cobra.Command{
	Use:     "gomd",
	Short:   asciiArt,
	Version: config.Version,
}

func init() {
	cmd.AddCommand(command.CLI)
}

func Run() error {
	return cmd.Execute()
}
