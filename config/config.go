package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/danilomarques1/fidus/clierror"
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
	cfg := &Config{}
	configLocation := cfg.getConfigLocation()
	cfg.configFolder = fmt.Sprintf("%v/%v", homeDir, configLocation)
	return cfg
}

func (cfg *Config) getConfigLocation() string {
	switch runtime.GOOS {
	case "windows":
		return "AppData/Local/fidus"
	default:
		return ".config/fidus"
	}
}

func (cfg *Config) CreateConfigFolder() error {
	if err := os.MkdirAll(cfg.configFolder, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetupLogger() error {
	logOutputFile, err := os.OpenFile(cfg.configFolder+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logOutputFile)
	return nil
}

func (cfg *Config) isTokenExpired(expiresAt int64) bool {
	tokenExpirationDate := time.UnixMilli(expiresAt * 1000)
	currentDate := time.Now()

	return tokenExpirationDate.Compare(currentDate) == -1
}

func (cfg *Config) GetToken() (string, error) {
	token, err := cfg.readToken()
	if err != nil {
		return "", clierror.ErrInvalidToken()
	}
	if cfg.isTokenExpired(token.ExpiresAt) {
		return "", clierror.ErrInvalidToken()
	}
	return token.AccessToken, nil
}

func (cfg *Config) RemoveToken() {
	if err := os.Truncate(fmt.Sprintf("%v/.token", cfg.configFolder), 0); err != nil {
		log.Printf("Error removing token %v", err)
	}
}

func (cfg *Config) readToken() (*Token, error) {
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
