package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gogatorconfig.json"

// TODO REMOVE THIS AFTER COMPLETING
const configFileName2 = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func configPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, configFileName), nil
}

// TODO : REMOVE THIS AFTER COMPLETION
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName2)
	return fullPath, nil
}

func (c *Config) Save() error {
	path, err := configPath()
	if err != nil {
		return err
	}
	// TODO : REMOVE THIS AFTER COMPLETION
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}
	// UNTIL HERE

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	path, err := configPath()
	if err != nil {
		return Config{}, nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return c.Save()
}
