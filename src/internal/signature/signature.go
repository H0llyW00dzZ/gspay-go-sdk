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

package signature

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"hash"
)

// Digest is a function that returns a new hash.Hash instance.
// Use this type to provide custom digest algorithms.
//
// Standard library hash functions can be used directly:
//   - crypto/md5.New (default)
//   - crypto/sha1.New
//   - crypto/sha256.New
//   - crypto/sha512.New
type Digest func() hash.Hash

// GenerateWithDigest creates a signature using the specified digest function.
// If digest is nil, MD5 is used as the default.
func GenerateWithDigest(data string, digest Digest) string {
	if digest == nil {
		// Default to MD5 for backward compatibility
		hash := md5.Sum([]byte(data))
		return hex.EncodeToString(hash[:])
	}

	h := digest()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Verify checks if the provided signature matches the expected signature.
// Uses constant-time comparison to prevent timing attacks.
func Verify(expected, actual string) bool {
	return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
}
