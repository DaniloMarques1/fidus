package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
	"github.com/danilomarques1/fidus/dto"
)

type StorePassword struct {
	passwordApi api.PasswordApi
	config      config.Config
}

func NewStorePassword() *StorePassword {
	cfg := config.NewConfig()
	passwordApi := api.NewPasswordApi()
	return &StorePassword{passwordApi: passwordApi, config: cfg}
}

func (storePassword *StorePassword) Execute(key, password string) error {
	body, err := dto.NewStorePasswordDto(key, password)
	if err != nil {
		return clierror.ErrInvalidParameters(err.Error())
	}
	token, err := storePassword.config.GetToken()
	if err != nil {
		return err
	}
	if err := storePassword.passwordApi.StorePassword(token, body); err != nil {
		return err
	}
	return nil
}
