package app

import (
	"errors"

	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/config"
)

type RetrievePassword struct {
	passwordApi api.PasswordApi
	config      *config.Config
}

func NewRetrievePassword() *RetrievePassword {
	passwordApi := api.NewPasswordApi("http://localhost:8080/fidus/password")
	config := config.NewConfig()
	return &RetrievePassword{passwordApi, config}
}

func (retrieve *RetrievePassword) Execute(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("Invalid parameters")
	}

	if !retrieve.config.IsTokenValid() {
		return "", errors.New("You need to authenticate again")
	}

	token, err := retrieve.config.ReadToken()
	if err != nil {
		return "", err
	}
	password, err := retrieve.passwordApi.RetrievePassword(token.AccessToken, key)
	if err != nil {
		return "", err
	}

	return password, nil
}
