package validate

import "testing"

func TestMasterPassword(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected bool
	}{
		{"Should return true", "1234@@Mock", true},
		{"Should return true", "1234!!moCK", true},
		{"Should return true", "@@@@@@12Mo", true},
		{"Should return false", "12345678", false},
		{"Should return false", "@@1234mock", false},
		{"Should return false", "1234455mockmockmock2131431", false},
		{"Should return false", "mockMock", false},
		{"Should return false", "1234Mo@", false},
		{"Should return false", "", false},
	}

	for _, tc := range cases {
		output := MasterPassword(tc.input)
		if output != tc.expected {
			t.Fatalf("Expected %v\n", tc.expected)
		}
	}
}
