package app

import (
	"testing"

	"github.com/danilomarques1/fidus/dto"
)

type masterApiMemory struct {
	users map[string]string
}

func newMasterApiMemory() *masterApiMemory {
	return &masterApiMemory{users: make(map[string]string)}
}

func (m *masterApiMemory) Register(body *dto.RegisterMasterDto) error {
	m.users[body.Email] = body.Password
	return nil
}

func (m *masterApiMemory) Authenticate(body *dto.AuthenticateMasterDto) (string, int64, error) {
	return "", 0, nil
}

func TestRegisterMaster(t *testing.T) {
	cases := []struct {
		label    string
		name     string
		email    string
		password string
		expected bool
	}{
		{"Should not return an error", "mock name", "mock@mail.com", "mock@@123Mock", false},
		{"Should not allow password length below 8", "mock name", "mock@mail.com", "mock", true},
		{"Should not allow wrong email type", "mock name", "mockcom", "mockpassword", true},
		{"Should not allow empty name", "", "mock@mail.com", "mock@@123Mock", true},
	}

	for _, tc := range cases {
		masterApi := &masterApiMemory{users: make(map[string]string)}
		registerMaster := RegisterMaster{masterApi}
		err := registerMaster.Execute(tc.name, tc.email, tc.password)
		out := err != nil
		if out != tc.expected {
			t.Fatalf("Should return error: %v\n", tc.expected)
		}
	}
}
