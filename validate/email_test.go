package validate

import "testing"

func TestEmail(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected bool
	}{
		{"It should return true", "mock@gmail.com", true},
		{"It should return false", "mock", false},
		{"It should return false", "mockgmailcom", false},
		{"It should return false", "mock@gmailcom", false},
		{"It should return false", "mock@emailcom", false},
		{"It should return true", "mock2345_231@gmail.com", true},
		{"It should return true", "mock2345_231@email.com", true},
		{"It should return true", "mock2345_231@hotmail.com", true},
		{"It should return true", "mock2345_231@outlook.com", true},
		{"It should return true", "mock2345_231@live.com.br", true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := Email(tc.input)
			if output != tc.expected {
				t.Error("Wrong output returned")
			}
		})
	}
}
