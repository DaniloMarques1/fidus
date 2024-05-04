package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/danilomarques1/fidus/dto"
)

type MasterApi interface {
	Register(dto.RegisterMasterDto) error
	Authenticate(dto.AuthenticateMasterDto) (string, int64, error)
}

type masterApi struct {
	baseUrl string
}

func NewMasterApi(baseUrl string) MasterApi {
	return &masterApi{baseUrl}
}

func (master *masterApi) Register(body dto.RegisterMasterDto) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := http.Post(master.baseUrl+"/register", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return readErrorMessageFromBody(resp.Body)
	}
	return nil
}

func (master *masterApi) Authenticate(body dto.AuthenticateMasterDto) (string, int64, error) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		return "", 0, err
	}
	resp, err := http.Post(master.baseUrl+"/authenticate", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println(err.Error())
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = readErrorMessageFromBody(resp.Body)
		log.Println(err.Error())
		return "", 0, err
	}
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return "", 0, err
	}
	respBody := &dto.AuthenticateMasterResponseDto{}
	if err := json.Unmarshal(responseBytes, respBody); err != nil {
		log.Println(err.Error())
		return "", 0, err
	}

	return respBody.AccessToken, respBody.ExpiresAt, nil
}

func readErrorMessageFromBody(body io.ReadCloser) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	respBody := &dto.ErrorResponseDto{}
	if err := json.Unmarshal(b, respBody); err != nil {
		return err
	}
	return errors.New(respBody.Message)
}
