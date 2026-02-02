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

package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNop(t *testing.T) {
	t.Run("implements Handler interface", func(t *testing.T) {
		var l Handler = Nop{}

		// Should not panic
		l.Debug("test", "key", "value")
		l.Info("test", "key", "value")
		l.Warn("test", "key", "value")
		l.Error("test", "key", "value")
	})
}

func TestStd(t *testing.T) {
	t.Run("logs at debug level", func(t *testing.T) {
		var buf bytes.Buffer
		l := NewStd(&buf, LevelDebug)

		l.Debug("debug message", "key", "value")
		l.Info("info message")
		l.Warn("warn message")
		l.Error("error message")

		output := buf.String()
		assert.Contains(t, output, "[DEBUG]")
		assert.Contains(t, output, "[INFO]")
		assert.Contains(t, output, "[WARN]")
		assert.Contains(t, output, "[ERROR]")
		assert.Contains(t, output, "key=value")
	})

	t.Run("respects log level", func(t *testing.T) {
		var buf bytes.Buffer
		l := NewStd(&buf, LevelWarn)

		l.Debug("debug message")
		l.Info("info message")
		l.Warn("warn message")
		l.Error("error message")

		output := buf.String()
		assert.NotContains(t, output, "[DEBUG]")
		assert.NotContains(t, output, "[INFO]")
		assert.Contains(t, output, "[WARN]")
		assert.Contains(t, output, "[ERROR]")
	})

	t.Run("handles odd number of key-value pairs", func(t *testing.T) {
		var buf bytes.Buffer
		l := NewStd(&buf, LevelDebug)

		l.Debug("test", "key1", "value1", "key2")

		output := buf.String()
		assert.Contains(t, output, "key1=value1")
		assert.Contains(t, output, "key2=(MISSING)")
	})

	t.Run("returns current level", func(t *testing.T) {
		l := NewStd(&bytes.Buffer{}, LevelInfo)
		assert.Equal(t, LevelInfo, l.Level())
	})
}

func TestDefault(t *testing.T) {
	t.Run("returns Std logger with debug level", func(t *testing.T) {
		l := Default()
		require.NotNil(t, l)
		assert.Equal(t, LevelDebug, l.Level())
	})
}

func TestFormatKeyValues(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			name:     "empty",
			input:    []any{},
			expected: "",
		},
		{
			name:     "single pair",
			input:    []any{"key", "value"},
			expected: "key=value",
		},
		{
			name:     "multiple pairs",
			input:    []any{"key1", "value1", "key2", 42},
			expected: "key1=value1 key2=42",
		},
		{
			name:     "odd number of args",
			input:    []any{"key1", "value1", "key2"},
			expected: "key1=value1 key2=(MISSING)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := FormatKeyValues(tc.input...)
			assert.Equal(t, tc.expected, result)
		})
	}
}
