package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mtyuksel/gum/models"
	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles map[string]models.Profile `yaml:"profiles"`
}

func ReadConfig() config {
	ex, err1 := os.Executable()

	if err1 != nil {
		panic(err1)
	}

	exPath := filepath.Dir(ex)
	yfile, err := ioutil.ReadFile(exPath + "/config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := config{}
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}

func GetProfileByName(profileName string) models.Profile {
	config := ReadConfig()

	profile := config.Profiles[profileName]
	fmt.Println(profile)

	return profile
}
