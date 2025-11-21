package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	b, err := os.ReadFile(".gogatorconfig.json")
	if err != nil {
		return Config{}
	}
	config := Config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return Config{}
	}
	return config
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(".gogatorconfig.json", data, 0o644)
	if err != nil {
		return err
	}
	return nil
}
