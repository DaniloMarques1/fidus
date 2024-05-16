package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/config"
)

type Keys struct {
	passwordApi api.PasswordApi
	config      config.Config
}

func NewKeys() *Keys {
	passwordApi := api.NewPasswordApi()
	cfg := config.NewConfig()
	return &Keys{passwordApi: passwordApi, config: cfg}
}

func (k *Keys) Execute() ([]string, error) {
	token, err := k.config.GetToken()
	if err != nil {
		return nil, err
	}
	keys, err := k.passwordApi.Keys(token)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
