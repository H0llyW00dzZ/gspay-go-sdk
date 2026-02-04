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

// Package payment provides payment services for the GSPAY2 SDK.
//
// This package supports creating and managing payments in multiple currencies
// including IDR (Indonesian Rupiah) and USDT (cryptocurrency).
//
// # IDR Payment Service
//
// Use [NewIDRService] to create IDR payments:
//
//	c := client.New("auth-key", "secret-key")
//	paymentSvc := payment.NewIDRService(c)
//
//	resp, err := paymentSvc.Create(ctx, &payment.IDRRequest{
//	    TransactionID: client.GenerateTransactionID("TXN"),
//	    Username:      "user123",
//	    Amount:        50000,
//	    Channel:       constants.ChannelQRIS,
//	})
//
// Available channels: QRIS, DANA, BNI Virtual Account.
//
// # USDT Payment Service
//
// Use [NewUSDTService] for cryptocurrency payments:
//
//	usdtSvc := payment.NewUSDTService(c)
//
//	resp, err := usdtSvc.Create(ctx, &payment.USDTRequest{
//	    TransactionID: client.GenerateTransactionID("USD"),
//	    Username:      "user123",
//	    Amount:        10.50,
//	})
//
// Note: USDT payments are not supported for Indonesian merchants
// due to government regulations.
//
// # Callback Verification
//
// Verify webhook callbacks with signature validation:
//
//	if err := paymentSvc.VerifyCallback(&callback); err != nil {
//	    // Invalid signature
//	}
//
// For additional security, verify the callback source IP:
//
//	if err := paymentSvc.VerifyCallbackWithIP(&callback, clientIP); err != nil {
//	    // Unauthorized IP or invalid signature
//	}
//
// # Payment Status
//
// Check payment status using:
//
//	status, err := paymentSvc.GetStatus(ctx, transactionID)
//	if status.Status.IsSuccess() {
//	    // Payment completed
//	}
package payment
