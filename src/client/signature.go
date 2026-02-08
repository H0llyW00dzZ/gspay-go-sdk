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

import "github.com/H0llyW00dzZ/gspay-go-sdk/src/internal/signature"

// GenerateSignature generates a signature for API requests.
// Uses the configured digest function (see [WithDigest]), or [crypto/md5] if not set.
func (c *Client) GenerateSignature(data string) string {
	return signature.GenerateWithDigest(data, c.digest)
}

// VerifySignature verifies a callback signature.
func (c *Client) VerifySignature(expected, actual string) bool {
	return signature.Verify(expected, actual)
}
