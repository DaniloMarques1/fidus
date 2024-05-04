package app

import (
	"errors"

	"github.com/danilomarques1/fidus/api"
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
	passwordApi := api.NewPasswordApi("http://localhost:8080/fidus/password")
	v := validate.Validate()
	return &StorePassword{passwordApi: passwordApi, validate: v, config: cfg}
}

func (storePassword *StorePassword) Execute(key, password string) error {
	body := &dto.StorePasswordDto{Key: key, Password: password}
	if err := storePassword.validate.Struct(body); err != nil {
		return errors.New("Invalid parameters")
	}
	if !storePassword.config.IsTokenValid() {
		return errors.New("Token expired. Please authenticate again")
	}
	token, err := storePassword.config.ReadToken()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := storePassword.passwordApi.StorePassword(token.AccessToken, body); err != nil {
		return err
	}
	return nil
}
