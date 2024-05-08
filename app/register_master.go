package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/dto"
	"github.com/danilomarques1/fidus/validate"
	"github.com/go-playground/validator/v10"
)

type RegisterMaster struct {
	masterApi api.MasterApi
	validate  *validator.Validate
}

func NewRegisterMaster() *RegisterMaster {
	masterApi := api.NewMasterApi()
	v := validate.Validate()
	return &RegisterMaster{masterApi: masterApi, validate: v}
}

func (master *RegisterMaster) Execute(name, email, password string) error {
	body := dto.RegisterMasterDto{Name: name, Email: email, Password: password}
	if err := master.validate.Struct(body); err != nil {
		return clierror.ErrInvalidParameters()
	}
	if err := master.masterApi.Register(body); err != nil {
		return err
	}
	return nil
}
