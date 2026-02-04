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

// Package constants provides constant values and types for the GSPAY2 API.
//
// This package defines all the constants, enums, and lookup functions needed
// for payment processing, including bank codes, payment channels, status codes,
// and API endpoints.
//
// # Payment Status
//
// The [PaymentStatus] type represents payment states:
//   - [StatusPending]: Payment is pending or expired (0)
//   - [StatusSuccess]: Payment completed successfully (1)
//   - [StatusFailed]: Payment failed (2)
//   - [StatusTimeout]: Payment timed out (4)
//
// Use helper methods to check status:
//
//	if status.IsSuccess() {
//	    // Payment completed
//	}
//
// # Payment Channels
//
// Available payment channels for IDR:
//   - [ChannelQRIS]: QRIS QR payment
//   - [ChannelDANA]: DANA e-wallet
//   - [ChannelBNI]: BNI Virtual Account
//
// # Bank Codes
//
// Bank code constants and lookup functions:
//   - [IsValidBankIDR]: Check if bank code is valid for IDR
//   - [IsValidBankMYR]: Check if bank code is valid for MYR
//   - [IsValidBankTHB]: Check if bank code is valid for THB
//   - [GetBankName]: Get bank name from code
//   - [GetBankCodes]: List all bank codes for a currency
//
// # API Endpoints
//
// Use [GetEndpoint] to retrieve API endpoint paths:
//
//	endpoint := constants.GetEndpoint(constants.EndpointIDRPayment)
package constants
