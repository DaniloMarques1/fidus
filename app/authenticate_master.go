package app

import (
	"errors"
	"log"

	"github.com/danilomarques1/fidus/app/api"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type AuthenticateMaster struct {
	masterApi api.MasterApi
	validate  *validator.Validate
}

func NewAuthenticateMaster() *AuthenticateMaster {
	masterApi := api.NewMasterApi("http://localhost:8080/fidus/master")
	v := validate.Validate()
	return &AuthenticateMaster{masterApi: masterApi, validate: v}
}

func (master *AuthenticateMaster) Execute(email, password string) error {
	body := dto.AuthenticateMasterDto{Email: email, Password: password}
	if err := master.validate.Struct(body); err != nil {
		return errors.New("Invalid parameters see help for usage")
	}
	accessToken, err := master.masterApi.Authenticate(body)
	if err != nil {
		return err
	}
	log.Printf("access token %v\n", accessToken)
	return nil
}
