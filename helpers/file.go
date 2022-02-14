package helpers

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"fmt"

	"github.com/mtyuksel/gum/models"
	"gopkg.in/yaml.v3"
)

type config struct {
	Profiles map[string]models.Profile `yaml:"profiles"`
}

func ReadConfig() (config, error) {
	file, err := filepath.Abs("config.yaml")
	if err != nil {
		return config{}, err
	}

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return config{}, err
	}

	data := config{}
	err = yaml.Unmarshal(yamlFile, &data)

	if err != nil {
		return config{}, err
	}

	return data, nil
}

func GetProfileByName(profileName string) (models.Profile, error) {
	config, err := ReadConfig()

	if err != nil {
		return models.Profile{}, err
	}

	profile := config.Profiles[profileName]

	if profile == (models.Profile{}) {
		return profile, errors.New("profile name does not exists please check config.yaml file")
	}

	return profile, nil
}

func ShowProfileNames() {
	config, err := ReadConfig()

	if err != nil {
		fmt.Println(err)
	}

	for key := range config.Profiles {
		fmt.Println(key)
	}
}
