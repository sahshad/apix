package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	BaseURL      string            `json:"base_url"`
	AuthToken    string            `json:"auth_token"`
	Environments map[string]string `json:"environments"`
	CurrentEnv   string            `json:"current_env"`
}

func LoadConfig() (*Config, error) {
	home, _ := os.UserHomeDir()
	file := home + "/.apixrc"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return &Config{}, nil
	}
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var cfg Config
	json.Unmarshal(b, &cfg)
	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	home, _ := os.UserHomeDir()
	file := home + "/.apixrc"
	b, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(file, b, 0644)
}
