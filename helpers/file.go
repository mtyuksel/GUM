package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mtyuksel/gum/models"
	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles map[string]models.Profile `yaml:"profiles"`
}

func ReadConfig() (config, error) {
	ex, err := os.Executable()

	if err != nil {
		return config{}, err
	}

	exPath := filepath.Dir(ex)
	yfile, err1 := ioutil.ReadFile(exPath + "/config.yaml")

	if err1 != nil {
		return config{}, err1
	}

	data := config{}
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		return config{}, err2
	}

	return data, nil
}

func GetProfileByName(profileName string) (models.Profile, error) {
	config, err := ReadConfig()

	if err != nil {
		return models.Profile{}, err
	}

	profile := config.Profiles[profileName]

	return profile, nil
}
