package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
)

type DeletePassword struct {
	passwordApi api.PasswordApi
	config      *config.Config
}

func NewDeletePassword() *DeletePassword {
	passwordApi := api.NewPasswordApi("http://localhost:8080/fidus/password")
	cfg := config.NewConfig()
	return &DeletePassword{passwordApi: passwordApi, config: cfg}
}

func (d *DeletePassword) Execute(key string) error {
	if len(key) == 0 {
		return clierror.ErrInvalidParameters()
	}
	token, err := d.config.GetToken()
	if err != nil {
		return err
	}
	if err := d.passwordApi.DeletePassword(token, key); err != nil {
		return err
	}

	return nil
}
