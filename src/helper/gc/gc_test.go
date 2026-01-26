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

package gc

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultPool(t *testing.T) {
	t.Run("gets and puts buffer", func(t *testing.T) {
		buf := Default.Get()
		require.NotNil(t, buf)
		assert.Equal(t, 0, buf.Len())

		// Put back
		Default.Put(buf)
	})

	t.Run("buffer write operations", func(t *testing.T) {
		buf := Default.Get()
		defer Default.Put(buf)

		// Test Write
		n, err := buf.Write([]byte("hello"))
		assert.NoError(t, err)
		assert.Equal(t, 5, n)
		assert.Equal(t, 5, buf.Len())
		assert.Equal(t, "hello", buf.String())

		// Test WriteString
		n, err = buf.WriteString(" world")
		assert.NoError(t, err)
		assert.Equal(t, 6, n)
		assert.Equal(t, 11, buf.Len())
		assert.Equal(t, "hello world", buf.String())

		// Test WriteByte
		err = buf.WriteByte('!')
		assert.NoError(t, err)
		assert.Equal(t, 12, buf.Len())
		assert.Equal(t, "hello world!", buf.String())

		// Test Bytes
		assert.Equal(t, []byte("hello world!"), buf.Bytes())
	})

	t.Run("buffer reset", func(t *testing.T) {
		buf := Default.Get()
		defer Default.Put(buf)

		buf.WriteString("test data")
		assert.Equal(t, 9, buf.Len())

		buf.Reset()
		assert.Equal(t, 0, buf.Len())
		assert.Equal(t, "", buf.String())
	})

	t.Run("buffer set operations", func(t *testing.T) {
		buf := Default.Get()
		defer Default.Put(buf)

		// Set with bytes
		buf.Set([]byte("bytes"))
		assert.Equal(t, "bytes", buf.String())

		// SetString
		buf.SetString("string")
		assert.Equal(t, "string", buf.String())
	})

	t.Run("buffer IO operations", func(t *testing.T) {
		buf := Default.Get()
		defer Default.Put(buf)

		// Test ReadFrom
		source := bytes.NewReader([]byte("read from"))
		n, err := buf.ReadFrom(source)
		assert.NoError(t, err)
		assert.Equal(t, int64(9), n)
		assert.Equal(t, "read from", buf.String())

		// Test WriteTo
		var dest bytes.Buffer
		n, err = buf.WriteTo(&dest)
		assert.NoError(t, err)
		assert.Equal(t, int64(9), n)
		assert.Equal(t, "read from", dest.String())
	})

	t.Run("buffer reuse after put", func(t *testing.T) {
		buf := Default.Get()
		buf.WriteString("first use")
		Default.Put(buf)

		// Get another buffer (may be the same one)
		buf2 := Default.Get()
		defer Default.Put(buf2)

		// Buffer should be clean
		assert.Equal(t, 0, buf2.Len())
		assert.Equal(t, "", buf2.String())
	})
}
