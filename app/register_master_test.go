package app

import (
	"errors"
	"testing"

	"github.com/danilomarques1/fidus/dto"
)

type masterApiMemory struct {
	users map[string]string
}

func newMasterApiMemory() *masterApiMemory {
	return &masterApiMemory{users: make(map[string]string)}
}

func (m *masterApiMemory) Register(body dto.RegisterMasterDto) error {
	m.users[body.Email] = body.Password
	return nil
}

func (m *masterApiMemory) Authenticate(body dto.AuthenticateMasterDto) error {
	password := m.users[body.Email]
	if password != body.Password {
		return errors.New("Unauthorized")
	}
	return nil
}

func TestRegisterMaster(t *testing.T) {
	masterApi := &masterApiMemory{users: make(map[string]string)}
	registerMaster := RegisterMaster{masterApi}
	err := registerMaster.Execute("Mock Name", "mock@mail.com", "mockpassword")
	if err != nil {
		t.Error(err)
	}

	authenticateMaster := AuthenticateMaster{masterApi}
	err = authenticateMaster.Execute("mock@mail.com", "mockpassword")
	if err != nil {
		t.Error(err)
	}
}
