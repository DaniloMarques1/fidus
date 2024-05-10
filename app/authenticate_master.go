package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/config"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type AuthenticateMaster struct {
	masterApi api.MasterApi
	validate  *validator.Validate
	config    *config.Config
}

func NewAuthenticateMaster() *AuthenticateMaster {
	masterApi := api.NewMasterApi()
	v := validate.Validate()
	config := config.NewConfig()
	return &AuthenticateMaster{masterApi: masterApi, validate: v, config: config}
}

func (master *AuthenticateMaster) Execute(email, password string) error {
	body, err := dto.NewAuthenticateMasterDto(email, password)
	if err != nil {
		return clierror.ErrInvalidParameters()
	}
	accessToken, expiresAt, err := master.masterApi.Authenticate(body)
	if err != nil {
		return err
	}
	if err := master.config.SaveToken(accessToken, expiresAt); err != nil {
		return err
	}
	return nil
}
