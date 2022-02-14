package helpers

import (
	"fmt"
	"os/exec"

	"github.com/mtyuksel/gum/models"
)

func SetProfile(profile models.Profile) error {
	//To change username
	cmdUserName := exec.Command("git", "config", "--global", "user.name", fmt.Sprintf("\"%s\"", profile.Username))
	err := cmdUserName.Run()

	if err != nil {
		return err
	}

	//To change email
	cmdEmail := exec.Command("git", "config", "--global", "user.email", fmt.Sprintf("\"%s\"", profile.Email))
	err = cmdEmail.Run()

	if err != nil {
		return err
	}

	return nil
}
