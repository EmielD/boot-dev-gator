package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("could not get users home directory: %v", err)
	}

	dat, err := os.ReadFile(userHomeDir + "/.gatorconfig.json")
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(dat, &config); err != nil {
		return Config{}, fmt.Errorf("could not decode configuration file: %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(name string) (Config, error) {
	c.Current_user_name = name

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("could not get users home directory: %v", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return Config{}, fmt.Errorf("could not encode user: %v", err)
	}

	err = os.WriteFile(userHomeDir+"/.gatorconfig.json", data, 0644)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %v", err)
	}

	return *c, nil
}
