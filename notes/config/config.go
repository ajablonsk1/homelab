package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configName = ".homelab.yaml"

type Config struct {
	NotesPath     string `yaml:"notes"`
	TemplatesPath string `yaml:"templates"`
}

func Get() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error with retrieving user home directory: %s", err)
	}

	content, err := os.ReadFile(homeDir + "/" + configName)
	if err != nil {
		log.Fatalf("Could not open config file: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Encountered error unmarshalling config file: %s", err)
	}

	return config
}
