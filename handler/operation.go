package handler

import (
	"fmt"
	"log"

	"github.com/mtyuksel/gum/helpers"
)

func SetProfile(profileName string) {
	profile, err := helpers.GetProfileByName(profileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = helpers.SetProfile(profile)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("-> Username: " + profile.Username)
	fmt.Println("-> Email   : " + profile.Email)
}

func ShowProfileNames() {
	helpers.ShowProfileNames()
}
