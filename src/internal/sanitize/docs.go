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
// This internal package ensures that sensitive information such as auth keys,
// account numbers, and account names are redacted or masked to prevent
// accidental exposure in logs.
//
// # Endpoint Sanitization
//
// The [Endpoint] function redacts auth keys from GSPAY2 API endpoint URLs.
// It handles both singular and plural operator path segments:
//
//	sanitize.Endpoint("/v2/integrations/operators/secret123/idr/payment")
//	// Returns: "/v2/integrations/operators/[REDACTED]/idr/payment"
//
//	sanitize.Endpoint("/v2/integrations/operator/secret123/balance")
//	// Returns: "/v2/integrations/operator/[REDACTED]/balance"
//
// # Account Data Masking
//
// The [AccountNumber] function masks account numbers, showing only the last
// 4 digits for identification while hiding the rest:
//
//	sanitize.AccountNumber("1234567890") // Returns: "****7890"
//	sanitize.AccountNumber("123")        // Returns: "****"
//
// The [AccountName] function masks account holder names, showing only the
// first character of each word:
//
//	sanitize.AccountName("John Doe") // Returns: "J*** D***"
//	sanitize.AccountName("Alice")    // Returns: "A***"
package sanitize
