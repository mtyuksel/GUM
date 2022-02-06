package main

import (
	"fmt"
	"os"

	"github.com/mtyuksel/gum/handler"
)

func main() {
	argsCount := len(os.Args)
	var operation string

	if argsCount > 1 {
		operation = os.Args[1]
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	if operation == "set" {
		if argsCount != 3 {
			fmt.Println("please enter a profile name to set")
			return
		}
		profileName := os.Args[2]
		handler.SetProfile(profileName)
		return
	}

	fmt.Println("invalid operation type")
}
