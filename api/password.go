package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/danilomarques1/fidus/dto"
)

type PasswordApi interface {
	StorePassword(token string, body *dto.StorePasswordDto) error
	RetrievePassword(token, key string) (string, error)
	DeletePassword(token, key string) error
	UpdatePassword(token, key string, body *dto.UpdatePasswordDto) error
}

type passwordApi struct {
	baseUrl string
}

func NewPasswordApi() PasswordApi {
	baseUrl := "http://localhost:8080/fidus/password"
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

func (p *passwordApi) RetrievePassword(token, key string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, p.baseUrl+"/retrieve", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	urlQuery := req.URL.Query()
	urlQuery.Add("key", key)
	req.URL.RawQuery = urlQuery.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", readErrorMessageFromBody(resp.Body)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody := &dto.RetrievePasswordDto{}
	if err := json.Unmarshal(b, respBody); err != nil {
		return "", err
	}

	return respBody.Password, nil
}

func (p *passwordApi) DeletePassword(token, key string) error {
	req, err := http.NewRequest(http.MethodDelete, p.baseUrl+"/delete", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	urlQuery := req.URL.Query()
	urlQuery.Add("key", key)
	req.URL.RawQuery = urlQuery.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return readErrorMessageFromBody(resp.Body)
	}

	return nil
}

func (p *passwordApi) UpdatePassword(token, key string, body *dto.UpdatePasswordDto) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, p.baseUrl+"/update", bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	q := req.URL.Query()
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
