package dto

type StorePasswordDto struct {
	Key      string `json:"key" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
