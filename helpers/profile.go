package helpers

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mtyuksel/gum/models"
)

func SetProfile(profile models.Profile) {
	//To change username
	cmdUserName := exec.Command("git", "config", "--global", "user.name", fmt.Sprintf("\"%s\"", profile.Username))
	err := cmdUserName.Run()

	if err != nil {
		log.Fatal(err)
	}

	//To change email
	cmdEmail := exec.Command("git", "config", "--global", "user.email", fmt.Sprintf("\"%s\"", profile.Email))
	err2 := cmdEmail.Run()

	if err2 != nil {
		log.Fatal(err2)
	}
}
