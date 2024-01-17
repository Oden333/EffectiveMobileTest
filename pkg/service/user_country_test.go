package service_test

import (
	"EMtest/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {

	type tc struct {
		name          string
		value         string
		expectedError error
		expectedValue string
	}
	tests := []tc{
		{
			name:          "Valid John",
			value:         "John",
			expectedError: nil,
			expectedValue: "IE",
		},
		{
			name:          "Invalid name",
			value:         "afskljh",
			expectedError: service.ErrorEmptyCountry,
			expectedValue: "",
		},
		{
			name:          "Empty name",
			value:         "",
			expectedError: service.ErrorEmptyName,
			expectedValue: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			country, err := service.GetCountry(test.value)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedValue, country)
		})
	}
}
