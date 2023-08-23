package main

import (
	"fmt"

	"github.com/JubaerHossain/gomd/command"
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
	Use:   "gomd",
	Short: asciiArt,
}

func AddCommand(cmd *cobra.Command) {
	cmd.AddCommand(cmd)
}

func init() {
	cmd.AddCommand(command.CLI)
}

// run is the main function to execute the gomd command
func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
