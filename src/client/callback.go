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

package client

import (
	"net"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/errors"
)

// VerifyCallbackIP verifies that the callback request originates from a whitelisted IP.
//
// Returns nil if the IP is whitelisted or if the whitelist is empty.
// Returns ErrIPNotWhitelisted if the IP is not in the whitelist.
// Returns ErrInvalidIPAddress if the IP address format is invalid.
func (c *Client) VerifyCallbackIP(ipStr string) error {
	// If no whitelist configured, skip IP validation
	if len(c.CallbackIPWhitelist) == 0 {
		return nil
	}

	// Strip port if present
	host := ipStr
	if h, _, err := net.SplitHostPort(ipStr); err == nil {
		host = h
	}

	// Validate IP format
	if net.ParseIP(host) == nil {
		return c.Error(errors.ErrInvalidIPAddress)
	}

	// Check whitelist
	if !c.IsIPWhitelisted(ipStr) {
		return c.Error(errors.ErrIPNotWhitelisted)
	}

	return nil
}
