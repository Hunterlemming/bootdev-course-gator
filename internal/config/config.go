package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	url, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(url)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %s", err)
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config: %s", err)
	}

	return cfg, nil
}

func (c *Config) SetUser(user string) error {
	newCfg := Config{c.DBURL, user}

	err := write(newCfg)
	if err != nil {
		return err
	}

	c.CurrentUserName = newCfg.CurrentUserName
	return nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error looking for home dir: %s", err)
	}

	return fmt.Sprintf("%s/%s", dir, configFileName), nil
}

func write(cfg Config) error {
	content, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %s", err)
	}

	url, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(url, content, 0666)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}

	return nil
}
