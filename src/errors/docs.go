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

// Package errors provides error types and error handling utilities for the GSPAY2 SDK.
//
// This package defines structured error types with i18n support for localized
// error messages in English and Indonesian.
//
// # Error Types
//
// The package provides several error types:
//   - [APIError]: Errors returned from the GSPAY2 API
//   - [ValidationError]: Client-side validation failures
//   - [LocalizedError]: Wrapper for errors with localized messages
//
// # Sentinel Errors
//
// Common validation errors are defined as sentinel errors:
//   - [ErrInvalidTransactionID]: Invalid or missing transaction ID
//   - [ErrInvalidUsername]: Invalid or missing username
//   - [ErrInvalidAmount]: Invalid payment amount
//   - [ErrInvalidAccountNumber]: Invalid bank account number
//   - [ErrInvalidBankCode]: Invalid or unsupported bank code
//   - [ErrInvalidSignature]: Signature verification failed
//   - [ErrRequestFailed]: HTTP request failed
//
// # Usage
//
// Check for specific errors:
//
//	if errors.Is(err, errors.ErrInvalidAmount) {
//	    // Handle invalid amount
//	}
//
// Get API error details:
//
//	if apiErr := errors.GetAPIError(err); apiErr != nil {
//	    log.Printf("API Error %d: %s", apiErr.Code, apiErr.Message)
//	}
//
// Create localized errors:
//
//	err := errors.New(i18n.Indonesian, errors.ErrInvalidAmount)
//	// Error message will be in Indonesian
//
// # Error Wrapping
//
// The [New] function supports error wrapping for proper error chains:
//
//	err := errors.New(lang, errors.ErrRequestFailed, originalErr)
//	// errors.Is(err, errors.ErrRequestFailed) returns true
//	// errors.Unwrap(err) returns originalErr
package errors
