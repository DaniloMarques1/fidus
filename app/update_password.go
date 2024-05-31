package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
	"github.com/danilomarques1/fidus/dto"
)

type UpdatePassword struct {
	passwordApi api.PasswordApi
	cfg         config.Config
}

func NewUpdatePassword() *UpdatePassword {
	passwordApi := api.NewPasswordApi()
	cfg := config.NewConfig()
	return &UpdatePassword{passwordApi: passwordApi, cfg: cfg}
}

func (u *UpdatePassword) Execute(key, newPassword string) error {
	if len(key) == 0 {
		return clierror.ErrInvalidParameters("Key must not be empty")
	}
	body, err := dto.NewUpdatePasswordDto(newPassword)
	if err != nil {
		return clierror.ErrInvalidParameters(err.Error())
	}
	token, err := u.cfg.GetToken()
	if err != nil {
		return err
	}
	if err := u.passwordApi.UpdatePassword(token, key, body); err != nil {
		return err
	}
	return nil
}
