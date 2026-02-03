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
	"strings"
	"unicode/utf8"
)

// Endpoint redacts sensitive information like auth keys from endpoint URLs.
//
// This function handles the GSPAY2 API endpoint patterns:
//   - /v2/integrations/operator/{authkey}/...  (singular - e.g., balance)
//   - /v2/integrations/operators/{authkey}/... (plural - e.g., USDT)
//
// The auth key following "operator" or "operators" is replaced with "[REDACTED]".
//
// Example:
//
//	sanitize.Endpoint("/v2/integrations/operators/secret123/idr/payment")
//	// Returns: "/v2/integrations/operators/[REDACTED]/idr/payment"
func Endpoint(endpoint string) string {
	parts := strings.Split(endpoint, "/")
	for i, part := range parts {
		// Check for operator/operators and ensure there is a next part to redact
		if (part == "operator" || part == "operators") && i+1 < len(parts) {
			// Check if the next part is not empty
			if len(parts[i+1]) > 0 {
				parts[i+1] = "[REDACTED]"
				return strings.Join(parts, "/")
			}
		}
	}
	return endpoint
}

// AccountNumber masks an account number for safe logging, showing only the last 4 digits.
//
// This function handles various account number formats:
//   - Short numbers (≤4 chars): Returns fully masked "****"
//   - Standard numbers: Returns "****" + last 4 digits
//   - Empty string: Returns empty string
//
// Example:
//
//	sanitize.AccountNumber("1234567890")
//	// Returns: "****7890"
//
//	sanitize.AccountNumber("123")
//	// Returns: "****"
func AccountNumber(accountNumber string) string {
	if accountNumber == "" {
		return ""
	}

	// Count runes for proper Unicode handling
	runeCount := utf8.RuneCountInString(accountNumber)

	// For short account numbers, return fully masked
	if runeCount <= 4 {
		return "****"
	}

	// Get last 4 characters (runes)
	runes := []rune(accountNumber)
	return "****" + string(runes[runeCount-4:])
}

// AccountName masks an account name for safe logging, showing only initials.
//
// This function handles various name formats:
//   - Single word: Returns first character + "***" (e.g., "John" -> "J***")
//   - Multiple words: Returns first char of each word (e.g., "John Doe" -> "J*** D***")
//   - Empty string: Returns empty string
//
// Example:
//
//	sanitize.AccountName("John Doe")
//	// Returns: "J*** D***"
//
//	sanitize.AccountName("Alice")
//	// Returns: "A***"
func AccountName(accountName string) string {
	if accountName == "" {
		return ""
	}

	words := strings.Fields(accountName)
	if len(words) == 0 {
		return ""
	}

	masked := make([]string, len(words))
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			masked[i] = string(runes[0]) + "***"
		}
	}

	return strings.Join(masked, " ")
}
