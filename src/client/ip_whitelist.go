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
)

// parseIPWhitelist parses the IP whitelist into net.IP and net.IPNet for efficient checking.
func (c *Client) parseIPWhitelist() {
	c.parsedIPNets = nil
	c.parsedIPs = nil

	for _, ipStr := range c.CallbackIPWhitelist {
		// Try parsing as CIDR first
		if _, ipNet, err := net.ParseCIDR(ipStr); err == nil {
			c.parsedIPNets = append(c.parsedIPNets, ipNet)
			continue
		}

		// Try parsing as individual IP
		if ip := net.ParseIP(ipStr); ip != nil {
			c.parsedIPs = append(c.parsedIPs, ip)
		}
	}
}

// IsIPWhitelisted checks if the given IP address is in the whitelist.
//
// Returns true if:
//   - The whitelist is empty (IP validation disabled)
//   - The IP matches an individual whitelisted IP
//   - The IP falls within a whitelisted CIDR range
//
// The ipStr parameter can include a port (e.g., "192.168.1.1:8080"),
// which will be automatically stripped before validation.
func (c *Client) IsIPWhitelisted(ipStr string) bool {
	// If no whitelist configured, allow all IPs
	if len(c.CallbackIPWhitelist) == 0 {
		return true
	}

	// Strip port if present (handles both IPv4 and IPv6)
	host := ipStr
	if h, _, err := net.SplitHostPort(ipStr); err == nil {
		host = h
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return false
	}

	// Check individual IPs
	for _, whitelistedIP := range c.parsedIPs {
		if whitelistedIP.Equal(ip) {
			return true
		}
	}

	// Check CIDR ranges
	for _, ipNet := range c.parsedIPNets {
		if ipNet.Contains(ip) {
			return true
		}
	}

	return false
}
