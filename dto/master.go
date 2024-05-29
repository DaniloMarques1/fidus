package dto

import (
	"errors"

	"github.com/danilomarques1/fidus/validate"
)

type RegisterMasterDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRegisterMasterDto(name, email, password string) (*RegisterMasterDto, error) {
	if len(name) == 0 {
		return nil, errors.New("Name is invalid")
	}
	if len(email) == 0 || !validate.Email(email) {
		return nil, errors.New("Email is invalid")
	}
	if !validate.MasterPassword(password) {
		return nil, errors.New("Password is invalid")
	}

	registerMasterDto := &RegisterMasterDto{Name: name, Email: email, Password: password}
	return registerMasterDto, nil
}

type AuthenticateMasterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthenticateMasterDto(email, password string) (*AuthenticateMasterDto, error) {
	if len(email) == 0 || !validate.Email(email) {
		return nil, errors.New("Email is invalid")
	}
	if !validate.MasterPassword(password) {
		return nil, errors.New("Password is invalid")
	}
	authenticateMasterDto := &AuthenticateMasterDto{Email: email, Password: password}
	return authenticateMasterDto, nil
}

type AuthenticateMasterResponseDto struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type ResetMasterPasswordDto struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// TODO add validations
func NewResetMasterPassword(email, oldPassword, newPassword string) (*ResetMasterPasswordDto, error) {
	return &ResetMasterPasswordDto{Email: email, OldPassword: oldPassword, NewPassword: newPassword}, nil
}
