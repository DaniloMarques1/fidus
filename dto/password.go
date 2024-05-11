package dto

import "errors"

type StorePasswordDto struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}

func NewStorePasswordDto(key, password string) (*StorePasswordDto, error) {
	if len(key) == 0 {
		return nil, errors.New("Key is invalid")
	}
	if len(password) == 0 {
		return nil, errors.New("Password is invalid")
	}
	storePasswordDto := &StorePasswordDto{Key: key, Password: password}
	return storePasswordDto, nil
}

type RetrievePasswordDto struct {
	MasterId string `json:"master_id"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

type UpdatePasswordDto struct {
	Password string `json:"password"`
}

func NewUpdatePasswordDto(password string) (*UpdatePasswordDto, error) {
	if len(password) == 0 {
		return nil, errors.New("Password is invalid")
	}
	updatePasswordDto := &UpdatePasswordDto{Password: password}
	return updatePasswordDto, nil
}
