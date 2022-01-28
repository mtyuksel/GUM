package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mtyuksel/gum/helpers"
)

func main() {
	profileName := os.Args[1]

	profile, err := helpers.GetProfileByName(profileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	err1 := helpers.SetProfile(profile)

	if err1 != nil {
		log.Fatal(err1)
		return
	}

	fmt.Println("-> Username: " + profile.Username)
	fmt.Println("-> Email   : " + profile.Email)
}
