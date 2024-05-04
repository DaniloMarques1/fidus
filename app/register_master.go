package app

import (
	"errors"

	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type RegisterMaster struct {
	masterApi api.MasterApi
	validate  *validator.Validate
}

func NewRegisterMaster() *RegisterMaster {
	masterApi := api.NewMasterApi("http://localhost:8080/fidus/master")
	v := validate.Validate()
	return &RegisterMaster{masterApi: masterApi, validate: v}
}

func (master *RegisterMaster) Execute(name, email, password string) error {
	body := dto.RegisterMasterDto{Name: name, Email: email, Password: password}
	if err := master.validate.Struct(body); err != nil {
		return errors.New("Invalid parameters see help for detail on usage")
	}
	if err := master.masterApi.Register(body); err != nil {
		return err
	}
	return nil
}
