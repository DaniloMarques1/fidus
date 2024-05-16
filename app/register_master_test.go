package app

import (
	"errors"
	"testing"

	"github.com/danilomarques1/fidus/dto"
)

type registerMasterApiMemory struct {
	shouldReturnError bool
}

func newRegisterMasterApiMemory(shouldReturnError bool) *registerMasterApiMemory {
	return &registerMasterApiMemory{shouldReturnError: shouldReturnError}
}

func (m *registerMasterApiMemory) Register(body *dto.RegisterMasterDto) error {
	if m.shouldReturnError {
		return errors.New("returning an error")
	}
	return nil
}

func (m *registerMasterApiMemory) Authenticate(body *dto.AuthenticateMasterDto) (string, int64, error) {
	return "", 0, nil
}

func TestRegisterMaster(t *testing.T) {
	cases := []struct {
		label         string
		name          string
		email         string
		password      string
		expectedError bool
	}{
		{"Should not return an error", "mock name", "mock@mail.com", "mock@@123Mock", false},
		{"Should not allow password length below 8", "mock name", "mock@mail.com", "mock", true},
		{"Should not allow wrong email type", "mock name", "mockcom", "mockpassword", true},
		{"Should not allow empty name", "", "mock@mail.com", "mock@@123Mock", true},
	}

	for _, tc := range cases {
		masterApi := newRegisterMasterApiMemory(false)
		registerMaster := RegisterMaster{masterApi}
		err := registerMaster.Execute(tc.name, tc.email, tc.password)
		out := err != nil
		if out != tc.expectedError {
			t.Fatalf("Should return error: %v\n", tc.expectedError)
		}
	}
}

func TestRegisterMasterError(t *testing.T) {
	masterApi := newRegisterMasterApiMemory(true)
	registerMaster := RegisterMaster{masterApi}
	err := registerMaster.Execute("Mock Name", "mock@mail.com", "mock123@@Mock")
	if err.Error() != "returning an error" {
		t.Fatalf("Wrong error returned %v\n", err)
	}
}
