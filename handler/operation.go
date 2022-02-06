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

	err1 := helpers.SetProfile(profile)

	if err1 != nil {
		log.Fatal(err1)
		return
	}

	fmt.Println("-> Username: " + profile.Username)
	fmt.Println("-> Email   : " + profile.Email)
}
