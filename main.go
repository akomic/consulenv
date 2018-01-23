package main

import (
	"fmt"
	"os"

	"consulenv/commands"
)

//
func main() {
	if err := commands.MyCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
