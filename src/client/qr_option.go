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

package client

import (
	"image/color"

	goqrcode "github.com/skip2/go-qrcode"
)

// QRRecoveryLevel represents the QR code error correction level.
//
// Higher levels allow more of the QR code to be damaged while remaining readable,
// but produce physically larger codes.
type QRRecoveryLevel = goqrcode.RecoveryLevel

// Error correction levels for QR code generation.
const (
	// QRRecoveryLow provides approximately 7% error recovery.
	QRRecoveryLow QRRecoveryLevel = goqrcode.Low
	// QRRecoveryMedium provides approximately 15% error recovery (default).
	QRRecoveryMedium QRRecoveryLevel = goqrcode.Medium
	// QRRecoveryHigh provides approximately 25% error recovery.
	QRRecoveryHigh QRRecoveryLevel = goqrcode.High
	// QRRecoveryHighest provides approximately 30% error recovery.
	QRRecoveryHighest QRRecoveryLevel = goqrcode.Highest
)

// Default values for QR code generation.
const (
	// DefaultQRSize is the default QR code image width and height in pixels.
	DefaultQRSize = 256
)

// qrConfig holds the QR code generation settings.
type qrConfig struct {
	size            int
	level           QRRecoveryLevel
	foregroundColor color.Color
	backgroundColor color.Color
}

// qrDefaults returns a qrConfig with default settings.
func qrDefaults() *qrConfig {
	return &qrConfig{
		size:            DefaultQRSize,
		level:           QRRecoveryMedium,
		foregroundColor: color.Black,
		backgroundColor: color.White,
	}
}

// QROption is a functional option for configuring QR code generation.
//
// Pass these to [WithQRCodeOptions] when creating a [Client].
type QROption func(*qrConfig)

// WithQRSize sets the QR code image size in pixels (width and height).
//
// A positive value sets a fixed image size.
// Default is 256 pixels.
//
// Example:
//
//	client.New("auth", "secret",
//	    client.WithQRCodeOptions(client.WithQRSize(512)),
//	)
func WithQRSize(pixels int) QROption {
	return func(c *qrConfig) {
		if pixels > 0 {
			c.size = pixels
		}
	}
}

// WithQRRecoveryLevel sets the QR code error correction level.
//
// Higher levels make the QR code more resilient to damage but increase size.
// Default is [QRRecoveryMedium].
//
// Example:
//
//	client.New("auth", "secret",
//	    client.WithQRCodeOptions(client.WithQRRecoveryLevel(client.QRRecoveryHigh)),
//	)
func WithQRRecoveryLevel(level QRRecoveryLevel) QROption {
	return func(c *qrConfig) {
		c.level = level
	}
}

// WithQRForegroundColor sets the foreground (module) color of the QR code.
//
// Default is [color.Black].
//
// Example:
//
//	client.WithQRForegroundColor(color.RGBA{R: 0, G: 0, B: 128, A: 255})
func WithQRForegroundColor(fg color.Color) QROption {
	return func(c *qrConfig) {
		if fg != nil {
			c.foregroundColor = fg
		}
	}
}

// WithQRBackgroundColor sets the background color of the QR code.
//
// Default is [color.White].
//
// Example:
//
//	client.WithQRBackgroundColor(color.RGBA{R: 240, G: 240, B: 240, A: 255})
func WithQRBackgroundColor(bg color.Color) QROption {
	return func(c *qrConfig) {
		if bg != nil {
			c.backgroundColor = bg
		}
	}
}
