package app

import (
	"github.com/danilomarques1/fidus/api"
	"github.com/danilomarques1/fidus/clierror"
	"github.com/danilomarques1/fidus/dto"
)

type ResetMasterPassword struct {
	masterApi api.MasterApi
}

func NewResetMasterPassword() *ResetMasterPassword {
	masterApi := api.NewMasterApi()
	return &ResetMasterPassword{masterApi: masterApi}
}

func (r *ResetMasterPassword) Execute(email, oldPassword, newPassword string) error {
	body, err := dto.NewResetMasterPasswordDto(email, oldPassword, newPassword)
	if err != nil {
		return clierror.ErrInvalidParameters(err.Error())
	}
	if err := r.masterApi.ResetMasterPassword(body); err != nil {
		return err
	}
	return nil
}
