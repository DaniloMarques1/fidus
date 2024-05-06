package dto

type StorePasswordDto struct {
	Key      string `json:"key" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RetrievePasswordDto struct {
	MasterId string `json:"master_id"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

type UpdatePasswordRequestDto struct {
	Password string `json:"password" validate:"required"`
}
