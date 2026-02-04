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

// Package sanitize provides utilities for sanitizing sensitive data
// from strings before logging or displaying to users.
//
// This internal package ensures that sensitive information such as auth keys
// and operator IDs are redacted from API endpoint URLs and error messages
// to prevent accidental exposure in logs.
//
// # Usage
//
// The [Endpoint] function sanitizes API URLs:
//
//	url := "https://api.example.com/operator/ABC123/payment"
//	sanitized := sanitize.Endpoint(url)
//	// Returns: "https://api.example.com/operator/[REDACTED]/payment"
//
// # Redacted Patterns
//
// The following patterns are sanitized:
//   - Auth keys in URL paths (e.g., /auth/{key}/...)
//   - Operator IDs in URL paths (e.g., /operator/{id}/... or /operators/{id}/...)
//   - Secret keys appearing in query strings
//
// All sensitive values are replaced with "[REDACTED]".
package sanitize
