package main

import (
	"fmt"
	"os"

	"github.com/mtyuksel/gum/helpers"
)

func main() {
	profileName := os.Args[1]

	profile := helpers.GetProfileByName(profileName)
	fmt.Println(profile.Email)
	fmt.Println(profile.Username)
}
