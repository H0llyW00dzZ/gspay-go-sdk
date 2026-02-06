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

// Generate creates an MD5 signature (lowercase hex string).
// This is a convenience function that uses MD5 as the digest.
//
// Deprecated: Use [GenerateWithDigest] instead.
func Generate(data string) string {
	return GenerateWithDigest(data, nil)
}
