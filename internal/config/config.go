package config

import (
	"os"
	"path/filepath"
	"encoding/json"
	"bufio"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DB_URL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath() 
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err 
	}
	defer file.Close()
	
	var cfg Config 
	err = json.NewDecoder(bufio.NewReader(file)).Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	return write(cfg)
}

func write(cfg *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0640)
	if err != nil {
		return err
	}
	
	return nil
}


func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}
