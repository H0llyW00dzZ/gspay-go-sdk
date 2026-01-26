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
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GenerateTransactionID generates a unique transaction ID suitable for GSPAY2 API.
//
// Format: PREFIX + YmdHis + random (total max 20 chars)
//
// Example:
//
//	txnID := client.GenerateTransactionID("TXN")
//	// Result: "TXN20260126143022123"
func GenerateTransactionID(prefix string) string {
	// Ensure prefix is short enough (max 3 chars)
	if len(prefix) > 3 {
		prefix = prefix[:3]
	}

	// YmdHis = 14 chars, prefix = 3 chars, random = 3 chars = 20 total
	timestamp := time.Now().Format("20060102150405")

	// Generate cryptographically secure random number
	randomNum, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		// Fallback to timestamp-based uniqueness if crypto/rand fails
		return fmt.Sprintf("%s%s%03d", prefix, timestamp, time.Now().Nanosecond()%1000)
	}

	return fmt.Sprintf("%s%s%03d", prefix, timestamp, randomNum.Int64())
}

// GenerateUUIDTransactionID generates a cryptographically secure UUID v4 transaction ID.
//
// This provides stronger uniqueness guarantees than the timestamp-based approach.
// The UUID is truncated to 20 characters to fit GSPAY2 API limits.
//
// Format: PREFIX + UUID (first 17 chars of UUID v4, total max 20 chars)
//
// Example:
//
//	txnID := client.GenerateUUIDTransactionID("TXN")
//	// Result: "TXN3d66c16c9db64210a"
func GenerateUUIDTransactionID(prefix string) string {
	// Ensure prefix is short enough (max 3 chars)
	if len(prefix) > 3 {
		prefix = prefix[:3]
	}

	// Generate cryptographically secure UUID v4
	id, err := uuid.NewRandom()
	if err != nil {
		// Fallback to timestamp-based ID if UUID generation fails
		return GenerateTransactionID(prefix)
	}

	// Convert to string and remove hyphens: "550e8400-e29b-41d4-a716-446655440000"
	uuidStr := strings.ReplaceAll(id.String(), "-", "")

	// Take first 17 characters to fit within 20-char limit (3 prefix + 17 uuid = 20)
	// This provides ~2^68 possible combinations (very high uniqueness)
	if len(uuidStr) > 17 {
		uuidStr = uuidStr[:17]
	}

	return fmt.Sprintf("%s%s", prefix, uuidStr)
}

// BuildReturnURL appends an encoded return URL to a payment URL for 1-way redirect.
//
// Example:
//
//	paymentURL := "https://pay.example.com/payment/123"
//	returnURL := "https://mysite.com/payment/complete"
//	fullURL := client.BuildReturnURL(paymentURL, returnURL)
//	// Result: "https://pay.example.com/payment/123?return=https%3A%2F%2Fmysite.com%2Fpayment%2Fcomplete"
func BuildReturnURL(paymentURL, returnURL string) string {
	separator := "?"
	if strings.Contains(paymentURL, "?") {
		separator = "&"
	}
	return paymentURL + separator + "return=" + url.QueryEscape(returnURL)
}

// FormatAmountIDR formats an integer amount as IDR currency string.
//
// Example:
//
//	formatted := client.FormatAmountIDR(50000)
//	// Result: "Rp 50.000"
func FormatAmountIDR(amount int64) string {
	// Convert to string and add thousand separators
	str := fmt.Sprintf("%d", amount)
	n := len(str)

	if n <= 3 {
		return "Rp " + str
	}

	// Add thousand separators
	var result strings.Builder
	result.WriteString("Rp ")

	remainder := n % 3
	if remainder > 0 {
		result.WriteString(str[:remainder])
		if n > 3 {
			result.WriteString(".")
		}
	}

	for i := remainder; i < n; i += 3 {
		if i > remainder {
			result.WriteString(".")
		}
		result.WriteString(str[i : i+3])
	}

	return result.String()
}

// FormatAmountUSDT formats a float amount as USDT currency string.
//
// Example:
//
//	formatted := client.FormatAmountUSDT(10.50)
//	// Result: "10.50 USDT"
func FormatAmountUSDT(amount float64) string {
	return fmt.Sprintf("%.2f USDT", amount)
}
