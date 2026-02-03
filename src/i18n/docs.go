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

// Package i18n provides internationalization support for the GSPAY2 SDK.
//
// This package handles localization of:
//   - Error messages (validation errors, API errors, sentinel errors)
//   - Log messages (request lifecycle, payment operations, retry logic)
//
// # Supported Languages
//
// Currently supported languages:
//   - English (default)
//   - Indonesian
//
// # Usage
//
// Set the language when creating a client:
//
//	c := client.New("auth-key", "secret-key",
//	    client.WithLanguage(i18n.Indonesian),
//	)
//
// Access translations directly:
//
//	msg := i18n.Get(i18n.Indonesian, i18n.MsgInvalidAmount)
//	// Returns: "jumlah pembayaran tidak valid"
//
// # Adding New Languages
//
// To add a new language:
//  1. Add a new Language constant in language.go
//  2. Update the IsValid() method to include the new language
//  3. Add translations for all MessageKey entries in the translations map in messages.go
//
// # Message Categories
//
// Messages are organized into categories:
//   - Sentinel error messages (MsgInvalidTransactionID, MsgInvalidAmount, etc.)
//   - Validation error messages (MsgMinAmountIDR, MsgMinAmountUSDT, etc.)
//   - Log messages for IDR Payment (LogCreatingIDRPayment, LogIDRPaymentCreated, etc.)
//   - Log messages for USDT Payment (LogCreatingUSDTPayment, LogUSDTPaymentCreated, etc.)
//   - Log messages for IDR Payout (LogCreatingIDRPayout, LogIDRPayoutCreated, etc.)
//   - Log messages for Balance (LogQueryingBalance, LogBalanceRetrieved)
//   - Log messages for HTTP Request (LogSendingRequest, LogRequestCompleted, etc.)
package i18n
