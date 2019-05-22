package config

import (
	"encoding/json"
	"os"

	"../models"
)

// LoadConfig : func
func LoadConfig() (configuration models.Config, errConfig error) {
	var config models.Config
	configurationFile, err := os.Open("conf.json")
	defer configurationFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configurationFile)
	jsonParser.Decode(&config)
	return config, nil
}
