package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type Config struct {
	configFolder string
}

var once sync.Once

func NewConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFolder := fmt.Sprintf("%v/.config_cache/fidus", homeDir)
	return &Config{configFolder}
}

func (cfg *Config) CreateConfigFolder() error {
	if err := os.MkdirAll(cfg.configFolder, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

// return true if there is a valid token stored
func (cfg *Config) IsTokenValid() bool {
	token, err := cfg.ReadToken()
	if err != nil {
		return false
	}
	return cfg.isTokenExpired(token.ExpiresAt)
}

func (cfg *Config) isTokenExpired(expiresAt int64) bool {
	t1 := time.UnixMilli(expiresAt * 1000)
	t2 := time.Now()

	return t1.Compare(t2) == +1
}

// TODO: need to improve this because i am reading the token
// to validate and then reading again to use it
func (cfg *Config) ReadToken() (*Token, error) {
	b, err := os.ReadFile(fmt.Sprintf("%v/.token", cfg.configFolder))
	if err != nil {
		return nil, err
	}
	token := &Token{}
	if err := json.Unmarshal(b, token); err != nil {
		return nil, err
	}
	return token, nil
}

func (cfg *Config) SaveToken(accessToken string, expiresAt int64) error {
	token := &Token{AccessToken: accessToken, ExpiresAt: expiresAt}
	file, err := os.Create(cfg.configFolder + "/.token")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(token); err != nil {
		return err
	}
	return nil
}
