package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

var configFileName = ".gatorconfig.json"

func Read() (config Config, err error) {

	configPath, err := getConfigPath()

	if err != nil {
		return Config{}, err
	}
	jsonFile, err := os.Open(configPath)

	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, err

}

func write(config Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(configPath)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(config)

}

func (config *Config) SetUser(user string) error {
	config.CurrentUserName = user
	return write(*config)
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, configFileName)
	return path, nil
}
