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

// Package amount provides utility functions for formatting monetary amounts.
//
// This package is primarily used for callback signature verification where
// amounts must be formatted with exactly 2 decimal places to match the
// expected signature format from GSPAY2.
//
// # Functions
//
// The package provides two formatting functions:
//
//   - [Format]: Parses a string amount and formats to 2 decimal places
//   - [FormatFloat]: Formats a float64 amount to 2 decimal places
//
// # Usage
//
// Format a string amount (e.g., from callback data):
//
//	formatted, err := amount.Format("10000", i18n.English)
//	// Returns: "10000.00", nil
//
// Format a float64 amount (e.g., for signature generation):
//
//	formatted := amount.FormatFloat(50000.5)
//	// Returns: "50000.50"
//
// # Precision Note
//
// The [Format] function uses float64 parsing which may have precision
// limitations for extremely large amounts (> 2^53). For typical payment
// amounts, this is not a concern.
package amount
