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

package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "redacts auth key in operators path",
			input:    "/v2/integrations/operators/secret123/idr/payment",
			expected: "/v2/integrations/operators/[REDACTED]/idr/payment",
		},
		{
			name:     "redacts auth key in operator path (singular)",
			input:    "/v2/integrations/operator/secret123/get/balance",
			expected: "/v2/integrations/operator/[REDACTED]/get/balance",
		},
		{
			name:     "preserves path without auth key pattern",
			input:    "/v1/other/endpoint",
			expected: "/v1/other/endpoint",
		},
		{
			name:     "preserves non-matching v2 paths",
			input:    "/v2/other/path/here",
			expected: "/v2/other/path/here",
		},
		{
			name:     "handles short paths",
			input:    "/v2/integrations",
			expected: "/v2/integrations",
		},
		{
			name:     "handles empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "handles path with query string",
			input:    "/v2/integrations/operators/secret123/idr/payment?foo=bar",
			expected: "/v2/integrations/operators/[REDACTED]/idr/payment?foo=bar",
		},
		{
			name:     "redacts auth key without leading slash",
			input:    "v2/integrations/operators/secret123/idr/payment",
			expected: "v2/integrations/operators/[REDACTED]/idr/payment",
		},
		{
			name:     "handles double slash before operator",
			input:    "site.com//v2/integrations/operators/secret123/idr/payment",
			expected: "site.com//v2/integrations/operators/[REDACTED]/idr/payment",
		},
		{
			name:     "handles operators at end of string (no key)",
			input:    "/v2/integrations/operators",
			expected: "/v2/integrations/operators",
		},
		{
			name:     "handles operators followed by empty segment",
			input:    "/v2/integrations/operators//something",
			expected: "/v2/integrations/operators//something",
		},
		{
			name:     "handles simple path without leading slash",
			input:    "operators/secret123/payment",
			expected: "operators/[REDACTED]/payment",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Endpoint(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAccountNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "masks standard account number",
			input:    "1234567890",
			expected: "****7890",
		},
		{
			name:     "masks long account number",
			input:    "12345678901234567890",
			expected: "****7890",
		},
		{
			name:     "masks exactly 5 digit account",
			input:    "12345",
			expected: "****2345",
		},
		{
			name:     "fully masks 4 digit account",
			input:    "1234",
			expected: "****",
		},
		{
			name:     "fully masks 3 digit account",
			input:    "123",
			expected: "****",
		},
		{
			name:     "fully masks 1 digit account",
			input:    "1",
			expected: "****",
		},
		{
			name:     "handles empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "handles account with letters",
			input:    "ABC1234567890",
			expected: "****7890",
		},
		{
			name:     "handles unicode characters",
			input:    "日本語1234567890",
			expected: "****7890",
		},
		{
			name:     "handles account with dashes",
			input:    "123-456-7890",
			expected: "****7890",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := AccountNumber(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestAccountName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "masks two word name",
			input:    "John Doe",
			expected: "J*** D***",
		},
		{
			name:     "masks single word name",
			input:    "Alice",
			expected: "A***",
		},
		{
			name:     "masks three word name",
			input:    "John Middle Doe",
			expected: "J*** M*** D***",
		},
		{
			name:     "handles empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "handles whitespace only",
			input:    "   ",
			expected: "",
		},
		{
			name:     "handles single character name",
			input:    "A",
			expected: "A***",
		},
		{
			name:     "handles unicode names",
			input:    "日本 太郎",
			expected: "日*** 太***",
		},
		{
			name:     "handles name with extra spaces",
			input:    "John   Doe",
			expected: "J*** D***",
		},
		{
			name:     "handles lowercase name",
			input:    "john doe",
			expected: "j*** d***",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := AccountName(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
