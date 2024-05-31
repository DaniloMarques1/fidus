package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/dto"
)

type RegisterMaster struct {
	masterApi api.MasterApi
}

func NewRegisterMaster() *RegisterMaster {
	masterApi := api.NewMasterApi()
	return &RegisterMaster{masterApi: masterApi}
}

func (master *RegisterMaster) Execute(name, email, password string) error {
	body, err := dto.NewRegisterMasterDto(name, email, password)
	if err != nil {
		return clierror.ErrInvalidParameters(err.Error())
	}
	if err := master.masterApi.Register(body); err != nil {
		return err
	}
	return nil
}
