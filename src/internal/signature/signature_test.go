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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("generates MD5 hash", func(t *testing.T) {
		data := "test data"
		sig := Generate(data)
		// MD5 of "test data" (without newline)
		expected := "eb733a00c0c9d336e65691a37ab54293"
		assert.Equal(t, expected, sig)
	})

	t.Run("generates lowercase hex", func(t *testing.T) {
		sig := Generate("hello")
		// Should be lowercase
		assert.Equal(t, "5d41402abc4b2a76b9719d911017c592", sig)
	})
}

func TestVerify(t *testing.T) {
	t.Run("returns true for matching signatures", func(t *testing.T) {
		expected := "9e107d9d372bb6826bd81d3542a419d6"
		actual := "9e107d9d372bb6826bd81d3542a419d6"
		assert.True(t, Verify(expected, actual))
	})

	t.Run("returns false for non-matching signatures", func(t *testing.T) {
		expected := "9e107d9d372bb6826bd81d3542a419d6"
		actual := "different"
		assert.False(t, Verify(expected, actual))
	})

	t.Run("returns false for different lengths", func(t *testing.T) {
		expected := "short"
		actual := "thisislongerthantheexpected"
		assert.False(t, Verify(expected, actual))
	})

	t.Run("handles empty strings", func(t *testing.T) {
		assert.True(t, Verify("", ""))
		assert.False(t, Verify("nonempty", ""))
		assert.False(t, Verify("", "nonempty"))
	})
}
