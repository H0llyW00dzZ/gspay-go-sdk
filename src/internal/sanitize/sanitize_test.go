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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Endpoint(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
