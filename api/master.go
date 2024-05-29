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
	Register(*dto.RegisterMasterDto) error
	Authenticate(*dto.AuthenticateMasterDto) (string, int64, error)
	ResetMasterPassword(*dto.ResetMasterPasswordDto) error
}

type masterApi struct {
	baseUrl string
}

func NewMasterApi() MasterApi {
	baseUrl := "https://fidusserver-5icrkm6i2q-uc.a.run.app/fidus/master"
	return &masterApi{baseUrl}
}

func (master *masterApi) Register(body *dto.RegisterMasterDto) error {
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

func (master *masterApi) Authenticate(body *dto.AuthenticateMasterDto) (string, int64, error) {
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

func (api *masterApi) ResetMasterPassword(body *dto.ResetMasterPasswordDto) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, api.baseUrl+"/reset/password", bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		err := readErrorMessageFromBody(resp.Body)
		log.Printf("Error %v\n", err)
		return err
	}
	return nil
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
