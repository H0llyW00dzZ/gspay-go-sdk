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

package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		key      EndpointKey
		expected string
	}{
		{
			name:     "EndpointBalance",
			key:      EndpointBalance,
			expected: "/v2/integrations/operator/%s/get/balance",
		},
		{
			name:     "EndpointIDRCreate",
			key:      EndpointIDRCreate,
			expected: "/v2/integrations/operators/%s/idr/payment",
		},
		{
			name:     "EndpointIDRStatus",
			key:      EndpointIDRStatus,
			expected: "/v2/integrations/operators/%s/idr/getpayment",
		},
		{
			name:     "EndpointUSDTCreate",
			key:      EndpointUSDTCreate,
			expected: "/v2/integrations/operators/%s/cryptocurrency/trc20/usdt",
		},
		{
			name:     "EndpointPayoutIDRCreate",
			key:      EndpointPayoutIDRCreate,
			expected: "/v2/integrations/operators/%s/idr/payout",
		},
		{
			name:     "EndpointPayoutIDRStatus",
			key:      EndpointPayoutIDRStatus,
			expected: "/v2/integrations/operators/%s/idr/payout/status",
		},
		{
			name:     "Unknown Endpoint",
			key:      "unknown_endpoint",
			expected: "unknown_endpoint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetEndpoint(tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}
