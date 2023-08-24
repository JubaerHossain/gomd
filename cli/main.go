package cli

import (
	"fmt"

	"github.com/JubaerHossain/gomd/cli/create"
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
var command = &cobra.Command{
	Use:   "gomd",
	Short: asciiArt,
}

func AddCommand(command *cobra.Command) {
	command.AddCommand(command)
}

func init() {
	command.AddCommand(create.Create)
}

// run is the main function to execute the gomd command
func Run() {
	err := command.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
