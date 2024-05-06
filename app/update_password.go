package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type UpdatePassword struct {
	passwordApi api.PasswordApi
	validate    *validator.Validate
	cfg         *config.Config
}

func NewUpdatePassword() *UpdatePassword {
	passwordApi := api.NewPasswordApi("http://localhost:8080/fidus/password")
	cfg := config.NewConfig()
	v := validate.Validate()
	return &UpdatePassword{passwordApi: passwordApi, cfg: cfg, validate: v}
}

func (u *UpdatePassword) Execute(key, newPassword string) error {
	if len(key) == 0 {
		return clierror.ErrInvalidParameters()
	}
	body := &dto.UpdatePasswordRequestDto{Password: newPassword}
	if err := u.validate.Struct(body); err != nil {
		return clierror.ErrInvalidParameters()
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
