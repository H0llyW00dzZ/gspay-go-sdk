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
	"bytes"
	"image/color"
	"os"
	"path/filepath"
	"testing"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/errors"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sampleQRIS is a fake QRIS-format payment string for testing.
// This follows the EMV QR Code TLV structure but uses entirely fictitious values.
const sampleQRIS = "00020101021226670016COM.FAKEBANK.WWW01189360000000000000000214000000000000000303UKE51440014ID.CO.QRIS.WWW0303UKE0215ID0000000000001520400005303360540850000.005802ID5916Fake Merchant Co6007JAKARTA61051234562070103A0163049999"

func newTestClient(opts ...Option) *Client {
	return New("test-auth", "test-secret", opts...)
}

func TestClient_QREncode(t *testing.T) {
	c := newTestClient()

	t.Run("encodes content to PNG bytes", func(t *testing.T) {
		png, err := c.QR().Encode("hello world")
		require.NoError(t, err)
		assert.NotEmpty(t, png)
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, png[:4])
	})

	t.Run("encodes QRIS payload", func(t *testing.T) {
		png, err := c.QR().Encode(sampleQRIS)
		require.NoError(t, err)
		assert.NotEmpty(t, png)
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, png[:4])
	})

	t.Run("returns sentinel error for empty content", func(t *testing.T) {
		_, err := c.QR().Encode("")
		require.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrEmptyQRContent)
	})

	t.Run("returns localized error for empty content (Indonesian)", func(t *testing.T) {
		idClient := newTestClient(WithLanguage(i18n.Indonesian))
		_, err := idClient.QR().Encode("")
		require.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrEmptyQRContent)
		assert.Contains(t, err.Error(), "konten kode QR tidak boleh kosong")
	})

	t.Run("returns ErrQREncodeFailed for content too long", func(t *testing.T) {
		// QR codes have a maximum capacity of ~2,953 bytes.
		// Generate content that exceeds this limit.
		longContent := string(make([]byte, 5000))
		_, err := c.QR().Encode(longContent)
		require.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrQREncodeFailed)
		assert.NotErrorIs(t, err, errors.ErrEmptyQRContent)
		// Verify the underlying goqrcode error is visible in the chain
		assert.Contains(t, err.Error(), "failed to encode QR code")
		t.Logf("Full error: %v", err)
	})

	t.Run("returns localized ErrQREncodeFailed (Indonesian)", func(t *testing.T) {
		idClient := newTestClient(WithLanguage(i18n.Indonesian))
		longContent := string(make([]byte, 5000))
		_, err := idClient.QR().Encode(longContent)
		require.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrQREncodeFailed)
		assert.Contains(t, err.Error(), "gagal mengenkode kode QR")
		// Verify the underlying goqrcode error is also in the chain
		t.Logf("Full error: %v", err)
	})

	t.Run("respects custom size option", func(t *testing.T) {
		small := newTestClient(WithQRCodeOptions(WithQRSize(64)))
		large := newTestClient(WithQRCodeOptions(WithQRSize(1024)))

		smallPNG, err := small.QR().Encode("test")
		require.NoError(t, err)

		largePNG, err := large.QR().Encode("test")
		require.NoError(t, err)

		assert.Greater(t, len(largePNG), len(smallPNG))
	})

	t.Run("respects recovery level option", func(t *testing.T) {
		low := newTestClient(WithQRCodeOptions(WithQRRecoveryLevel(QRRecoveryLow)))
		high := newTestClient(WithQRCodeOptions(WithQRRecoveryLevel(QRRecoveryHighest)))

		lowPNG, err := low.QR().Encode("test")
		require.NoError(t, err)

		highPNG, err := high.QR().Encode("test")
		require.NoError(t, err)

		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, lowPNG[:4])
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, highPNG[:4])
	})

	t.Run("respects color options", func(t *testing.T) {
		c := newTestClient(WithQRCodeOptions(
			WithQRForegroundColor(color.RGBA{R: 0, G: 0, B: 128, A: 255}),
			WithQRBackgroundColor(color.RGBA{R: 240, G: 240, B: 240, A: 255}),
		))
		png, err := c.QR().Encode("test")
		require.NoError(t, err)
		assert.NotEmpty(t, png)
	})

	t.Run("ignores nil color options", func(t *testing.T) {
		c := newTestClient(WithQRCodeOptions(
			WithQRForegroundColor(nil),
			WithQRBackgroundColor(nil),
		))
		png, err := c.QR().Encode("test")
		require.NoError(t, err)
		assert.NotEmpty(t, png)
	})

	t.Run("ignores non-positive size", func(t *testing.T) {
		c := newTestClient(WithQRCodeOptions(WithQRSize(0)))
		png, err := c.QR().Encode("test")
		require.NoError(t, err)
		assert.NotEmpty(t, png)

		c2 := newTestClient(WithQRCodeOptions(WithQRSize(-1)))
		png2, err := c2.QR().Encode("test")
		require.NoError(t, err)
		assert.NotEmpty(t, png2)
	})
}

func TestClient_QRWrite(t *testing.T) {
	c := newTestClient()

	t.Run("writes PNG to buffer", func(t *testing.T) {
		var buf bytes.Buffer
		err := c.QR().Write(&buf, "hello world")
		require.NoError(t, err)
		assert.NotEmpty(t, buf.Bytes())
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, buf.Bytes()[:4])
	})

	t.Run("writes QRIS payload to buffer", func(t *testing.T) {
		large := newTestClient(WithQRCodeOptions(WithQRSize(512)))
		var buf bytes.Buffer
		err := large.QR().Write(&buf, sampleQRIS)
		require.NoError(t, err)
		assert.NotEmpty(t, buf.Bytes())
	})

	t.Run("returns error for empty content", func(t *testing.T) {
		var buf bytes.Buffer
		err := c.QR().Write(&buf, "")
		require.Error(t, err)
	})
}

func TestClient_QRWriteFile(t *testing.T) {
	c := newTestClient()

	t.Run("writes PNG to file", func(t *testing.T) {
		tmpDir := t.TempDir()
		filename := filepath.Join(tmpDir, "test_qr.png")

		err := c.QR().WriteFile(filename, "hello world")
		require.NoError(t, err)

		data, err := os.ReadFile(filename)
		require.NoError(t, err)
		assert.NotEmpty(t, data)
		assert.Equal(t, []byte{0x89, 0x50, 0x4E, 0x47}, data[:4])
	})

	t.Run("writes QRIS payload to file with options", func(t *testing.T) {
		high := newTestClient(WithQRCodeOptions(WithQRSize(512), WithQRRecoveryLevel(QRRecoveryHigh)))
		tmpDir := t.TempDir()
		filename := filepath.Join(tmpDir, "qris_payment.png")

		err := high.QR().WriteFile(filename, sampleQRIS)
		require.NoError(t, err)

		data, err := os.ReadFile(filename)
		require.NoError(t, err)
		assert.NotEmpty(t, data)
	})

	t.Run("returns error for empty content", func(t *testing.T) {
		tmpDir := t.TempDir()
		filename := filepath.Join(tmpDir, "empty.png")

		err := c.QR().WriteFile(filename, "")
		require.Error(t, err)

		_, err = os.Stat(filename)
		assert.True(t, os.IsNotExist(err))
	})
}

func TestQRDefaults(t *testing.T) {
	cfg := qrDefaults()

	assert.Equal(t, DefaultQRSize, cfg.size)
	assert.Equal(t, QRRecoveryMedium, cfg.level)
	assert.Equal(t, color.Black, cfg.foregroundColor)
	assert.Equal(t, color.White, cfg.backgroundColor)
}
