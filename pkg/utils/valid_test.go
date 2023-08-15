package utils

import (
	"testing"
)

func TestIsValidName(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
	}{
		{name: "John", expected: true},
		{name: "Mary", expected: true},
		{name: "John123", expected: false},
		{name: "Mary Doe", expected: false},
		{name: "john", expected: true},
		{name: "MARY", expected: true},
		{name: "John@Doe", expected: false},
	}

	for _, testCase := range testCases {
		actual := IsValidName(testCase.name)
		if actual != testCase.expected {
			t.Errorf("For name '%s', expected %v, but got %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{email: "test@example.com", expected: true},
		{email: "username@example.co.uk", expected: true},
		{email: "12345@example.com", expected: true},
		{email: "test@localhost", expected: false},
		{email: "test@example", expected: false},
		{email: "@example.com", expected: false},
	}

	for _, tc := range testCases {
		result := IsValidEmail(tc.email)
		if result != tc.expected {
			t.Errorf("Expected IsValidEmail(%q) to be %v, but got %v", tc.email, tc.expected, result)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{password: "Abc123!@#", expected: true},
		{password: "abcdefgh", expected: false},
		{password: "12345678", expected: false},
		{password: "ABCDEF", expected: false},
		{password: "abc 123", expected: false},
		{password: "abc123", expected: false},
		{password: "abcABC123", expected: false},
		{password: "abcABC123!", expected: true},
		{password: " abcABC123! ", expected: false},
	}

	for _, test := range tests {
		result := IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("Expected IsValidPassword(%q) to be %v, got %v", test.password, test.expected, result)
		}
	}
}
