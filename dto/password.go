package dto

type StorePasswordDto struct {
	Key      string `json:"key" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RetrievePasswordDto struct {
	MasterId string `json:"master_id"`
	Key      string `json:"key"`
	Password string `json:"password"`
}
