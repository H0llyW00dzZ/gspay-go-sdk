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

// Package client provides the HTTP client for the GSPAY2 API.
//
// This package is the foundation of the SDK, providing configuration,
// HTTP request handling, and utility functions for all API operations.
//
// # Creating a Client
//
// Use [New] with functional options pattern:
//
//	c := client.New("auth-key", "secret-key",
//	    client.WithTimeout(60*time.Second),
//	    client.WithRetries(5),
//	    client.WithLanguage(i18n.Indonesian),
//	)
//
// # Configuration Options
//
// Available options:
//   - [WithBaseURL]: Set custom API base URL
//   - [WithTimeout]: Set request timeout (default: 30s)
//   - [WithRetries]: Set retry attempts (default: 3)
//   - [WithRetryWait]: Set min/max wait between retries
//   - [WithHTTPClient]: Use custom http.Client
//   - [WithLanguage]: Set language for error and log messages
//   - [WithDebug]: Enable debug logging to stderr
//   - [WithLogger]: Set custom structured logger
//   - [WithDigest]: Set custom hash function for signatures (default: MD5)
//   - [WithCallbackIPWhitelist]: Set allowed IPs for callback verification
//   - [WithQRCodeOptions]: Configure QR code generation (size, recovery level, colors)
//
// # Retry Logic
//
// The client includes automatic retry with exponential backoff and jitter
// for transient failures (5xx errors, timeouts, connection issues).
//
// # Helper Functions
//
// Utility functions for common operations:
//   - [GenerateTransactionID]: Generate unique transaction IDs
//   - [GenerateUUIDTransactionID]: Generate UUID-based transaction IDs
//   - [BuildReturnURL]: Append return URL to payment URL
//   - [FormatAmountIDR]: Format IDR amounts (e.g., "Rp 50.000")
//   - [FormatAmountUSDT]: Format USDT amounts (e.g., "10.50 USDT")
//
// # Logging
//
// By default, logging is disabled. Enable it with:
//
//	// Quick debug mode
//	client.WithDebug(true)
//
//	// Custom logger
//	client.WithLogger(logger.NewStd(os.Stdout, logger.LevelInfo))
//
// See the [logger] subpackage for more logging options.
package client
