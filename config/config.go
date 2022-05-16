package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	CommandFile struct {
		Path     string `json:"path"`
		Filename string `json:"filename"`
	} `json:"command_file"`
	Port string `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
