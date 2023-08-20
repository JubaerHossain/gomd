package main

import (
	"fmt"

	"github.com/JubaerHossain/gomd/gomd"
)



// run is the main function to execute the gomd command
func main() {
	err := gomd.Run()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
