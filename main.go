package main

import (
	"fmt"
	"log"

	"github.com/JubaerHossain/gomd/config"
	"github.com/spf13/cobra"
)

var asciiArt = `
     ____ _____  ____ ___  ____/ /
    / __ / __ \/ __  __ \/ __  / 
   / /_/ / /_/ / / / / / / /_/ /  
   \__, /\____/_/ /_/ /_/\__,_/   
  /____/                          
`


// Create a Cobra command for gomd
var cmd = &cobra.Command{
	Use:   "gomd",
	Short: asciiArt,
	Long:  "gomd is a sleek and powerful markdown server designed to elegantly serve markdown files over HTTP.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(asciiArt)
		fmt.Println("✨ A CLI tool for building golang restful API application. ✨")
	},
	Version: config.Version,

}

// run is the main function to execute the gomd command
func main() {
	// Execute the Cobra command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
