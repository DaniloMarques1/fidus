package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type StorePassword struct {
	passwordApi api.PasswordApi
	validate    *validator.Validate
	config      *config.Config
}

func NewStorePassword() *StorePassword {
	cfg := config.NewConfig()
	passwordApi := api.NewPasswordApi()
	v := validate.Validate()
	return &StorePassword{passwordApi: passwordApi, validate: v, config: cfg}
}

func (storePassword *StorePassword) Execute(key, password string) error {
	body, err := dto.NewStorePasswordDto(key, password)
	if err != nil {
		return clierror.ErrInvalidParameters()
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
