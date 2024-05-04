package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/danilomarques1/fidus/dto"
)

type PasswordApi interface {
	StorePassword(token string, body *dto.StorePasswordDto) error
}

type passwordApi struct {
	baseUrl string
}

func NewPasswordApi(baseUrl string) PasswordApi {
	return &passwordApi{baseUrl}
}

func (p *passwordApi) StorePassword(token string, body *dto.StorePasswordDto) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, p.baseUrl+"/store", bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return readErrorMessageFromBody(res.Body)
	}

	return nil
}
