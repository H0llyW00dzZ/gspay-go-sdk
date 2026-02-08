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
	"io"
	"os"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/errors"
	goqrcode "github.com/skip2/go-qrcode"
)

// QR returns a handle to the QR code.
func (c *Client) QR() *QR {
	return &QR{client: c}
}

// QR provides QR code generation.
type QR struct{ client *Client }

// Encode encodes content into a QR code and returns the PNG image as bytes.
//
// The content is typically a QRIS payment string from the QR field
// of a payment response (e.g., payment.IDRResponse.QR).
// QR code settings are configured once via [WithQRCodeOptions].
//
// Example:
//
//	png, err := c.QR().Encode(resp.QR)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (q *QR) Encode(content string) ([]byte, error) {
	if content == "" {
		return nil, q.client.Error(errors.ErrEmptyQRContent)
	}

	qr, err := goqrcode.New(content, q.client.qrCfg.level)
	if err != nil {
		return nil, q.client.Error(errors.ErrEmptyQRContent, err)
	}

	qr.ForegroundColor = q.client.qrCfg.foregroundColor
	qr.BackgroundColor = q.client.qrCfg.backgroundColor

	return qr.PNG(q.client.qrCfg.size)
}

// Write encodes content into a QR code and writes the PNG image to w.
//
// Example:
//
//	w.Header().Set("Content-Type", "image/png")
//	err := c.QR().Write(w, resp.QR)
func (q *QR) Write(w io.Writer, content string) error {
	png, err := q.Encode(content)
	if err != nil {
		return err
	}

	_, err = w.Write(png)
	return err
}

// WriteFile encodes content into a QR code and saves the PNG image to a file.
//
// The file is created with permissions 0644. If the file already exists, it is overwritten.
//
// Example:
//
//	err := c.QR().WriteFile("payment_qr.png", resp.QR)
func (q *QR) WriteFile(filename, content string) error {
	png, err := q.Encode(content)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, png, 0644)
}
