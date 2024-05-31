package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
)

type RetrievePassword struct {
	passwordApi api.PasswordApi
	config      config.Config
}

func NewRetrievePassword() *RetrievePassword {
	passwordApi := api.NewPasswordApi()
	config := config.NewConfig()
	return &RetrievePassword{passwordApi, config}
}

func (retrieve *RetrievePassword) Execute(key string) (string, error) {
	if len(key) == 0 {
		return "", clierror.ErrInvalidParameters("Key must not be empty")
	}
	token, err := retrieve.config.GetToken()
	if err != nil {
		return "", err
	}
	password, err := retrieve.passwordApi.RetrievePassword(token, key)
	if err != nil {
		return "", err
	}

	return password, nil
}
