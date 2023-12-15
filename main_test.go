package main

import (
	"errors"
	"testing"
)

func TestValidateURL(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		expected error
	}{
		{"Valid URL", "https://example.com", nil},
		{"Empty URL", "", errors.New("url is required")},
		{"Missing Host", "https://", errors.New("missing url host")},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := validateURL(tc.input)
			if err != nil {
				if err.Error() != tc.expected.Error() {
					t.Errorf("expect error %v, but got %v", tc.expected, err)
				}
			} else {
				if err != tc.expected {
					t.Errorf("expect error %v, but got %v", tc.expected, err)
				}
			}
		})
	}
}
