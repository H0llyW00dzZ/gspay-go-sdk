// Copyright 2026 H0llyW00dzZ
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	t.Run("formats error without endpoint", func(t *testing.T) {
		err := &APIError{
			Code:    400,
			Message: "bad request",
		}
		expected := "gspay: API error 400: bad request"
		assert.Equal(t, expected, err.Error())
	})

	t.Run("formats error with endpoint", func(t *testing.T) {
		err := &APIError{
			Code:     500,
			Message:  "internal server error",
			Endpoint: "/api/test",
		}
		expected := "gspay: API error 500 on /api/test: internal server error"
		assert.Equal(t, expected, err.Error())
	})

	t.Run("sanitizes auth key in endpoint", func(t *testing.T) {
		err := &APIError{
			Code:     401,
			Message:  "unauthorized",
			Endpoint: "/v2/integrations/operators/secretkey123/idr/payment",
		}
		expected := "gspay: API error 401 on /v2/integrations/operators/[REDACTED]/idr/payment: unauthorized"
		assert.Equal(t, expected, err.Error())
	})
}

func TestIsAPIError(t *testing.T) {
	t.Run("returns true for APIError", func(t *testing.T) {
		err := &APIError{Code: 400, Message: "test"}
		assert.True(t, IsAPIError(err))
	})

	t.Run("returns false for other errors", func(t *testing.T) {
		err := errors.New("regular error")
		assert.False(t, IsAPIError(err))
	})

	t.Run("returns false for nil", func(t *testing.T) {
		assert.False(t, IsAPIError(nil))
	})
}

func TestGetAPIError(t *testing.T) {
	t.Run("extracts APIError", func(t *testing.T) {
		original := &APIError{Code: 404, Message: "not found"}
		wrapped := GetAPIError(original)
		assert.Equal(t, original, wrapped)
	})

	t.Run("returns nil for non-APIError", func(t *testing.T) {
		err := errors.New("regular error")
		assert.Nil(t, GetAPIError(err))
	})
}

func TestValidationError_Error(t *testing.T) {
	err := &ValidationError{
		Field:   "amount",
		Message: "must be positive",
	}
	expected := "gspay: validation error for amount: must be positive"
	assert.Equal(t, expected, err.Error())
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("bank_code", "invalid code")
	assert.Equal(t, "bank_code", err.Field)
	assert.Equal(t, "invalid code", err.Message)
}

func TestIsValidationError(t *testing.T) {
	t.Run("returns true for ValidationError", func(t *testing.T) {
		err := &ValidationError{Field: "test", Message: "error"}
		assert.True(t, IsValidationError(err))
	})

	t.Run("returns false for other errors", func(t *testing.T) {
		err := errors.New("regular error")
		assert.False(t, IsValidationError(err))
	})
}

func TestGetValidationError(t *testing.T) {
	t.Run("extracts ValidationError", func(t *testing.T) {
		original := &ValidationError{Field: "test", Message: "error"}
		extracted := GetValidationError(original)
		assert.Equal(t, original, extracted)
	})

	t.Run("returns nil for non-ValidationError", func(t *testing.T) {
		err := errors.New("regular error")
		assert.Nil(t, GetValidationError(err))
	})
}

func TestSentinelErrors(t *testing.T) {
	testCases := []struct {
		name string
		err  error
	}{
		{"ErrInvalidTransactionID", ErrInvalidTransactionID},
		{"ErrInvalidAmount", ErrInvalidAmount},
		{"ErrInvalidBankCode", ErrInvalidBankCode},
		{"ErrInvalidSignature", ErrInvalidSignature},
		{"ErrMissingCallbackField", ErrMissingCallbackField},
		{"ErrEmptyResponse", ErrEmptyResponse},
		{"ErrInvalidJSON", ErrInvalidJSON},
		{"ErrRequestFailed", ErrRequestFailed},
		{"ErrIPNotWhitelisted", ErrIPNotWhitelisted},
		{"ErrInvalidIPAddress", ErrInvalidIPAddress},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Error(t, tc.err)
			assert.NotEmpty(t, tc.err.Error())
		})
	}
}
