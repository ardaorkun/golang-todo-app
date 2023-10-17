package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string `yaml:"database_url"`
	Port        string `yaml:"port"`
}

func GetConfig() Config {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}
