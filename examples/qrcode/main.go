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

// Package main demonstrates QR code generation with the GSPAY Go SDK.
//
// This example shows how to:
//   - Configure QR code options (size, recovery level, colors)
//   - Create an IDR payment with QRIS channel
//   - Encode the QR response into a PNG image
//   - Save the QR code to a file
package main

import (
	"context"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/constants"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/errors"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/payment"
)

func main() {
	// Get credentials from environment variables
	authKey := os.Getenv("GSPAY_AUTH_KEY")
	secretKey := os.Getenv("GSPAY_SECRET_KEY")

	if authKey == "" || secretKey == "" {
		log.Fatal("Please set GSPAY_AUTH_KEY and GSPAY_SECRET_KEY environment variables")
	}

	// Create client with QR code options
	c := client.New(
		authKey,
		secretKey,
		client.WithTimeout(60*time.Second),
		client.WithQRCodeOptions(
			client.WithQRSize(512),
			client.WithQRRecoveryLevel(client.QRRecoveryHigh),
			client.WithQRForegroundColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}),
			client.WithQRBackgroundColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
		),
	)

	ctx := context.Background()

	// Create IDR payment with QRIS channel
	paymentSvc := payment.NewIDRService(c)

	fmt.Println("=== Creating IDR Payment (QRIS) ===")
	resp, err := paymentSvc.Create(ctx, &payment.IDRRequest{
		TransactionID: client.GenerateUUIDTransactionID("QR"),
		Username:      "demo_user",
		Amount:        50000,
		Channel:       constants.ChannelQRIS,
	})
	if err != nil {
		if apiErr := errors.GetAPIError(err); apiErr != nil {
			log.Fatalf("API Error: %d - %s", apiErr.Code, apiErr.Message)
		}
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Payment ID: %s\n", resp.IDRPaymentID)
	fmt.Printf("Amount: %s\n", client.FormatAmountIDR(50000))
	fmt.Printf("Payment URL: %s\n", resp.PaymentURL)
	fmt.Printf("Expires: %s\n", resp.ExpireDate)

	// Check if QR data is available
	if resp.QR == "" {
		log.Fatal("No QR code data in response")
	}

	fmt.Printf("QR Data Length: %d bytes\n", len(resp.QR))

	// Save QR code to file
	fmt.Println("\n=== Saving QR Code ===")
	filename := "payment_qr.png"
	if err := c.QR().WriteFile(filename, resp.QR); err != nil {
		log.Fatalf("Failed to save QR code: %v", err)
	}

	fmt.Printf("QR code saved to: %s\n", filename)
}
