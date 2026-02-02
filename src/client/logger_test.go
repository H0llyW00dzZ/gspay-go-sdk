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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockLogger captures log calls for testing.
type MockLogger struct {
	DebugCalls []logCall
	InfoCalls  []logCall
	WarnCalls  []logCall
	ErrorCalls []logCall
}

type logCall struct {
	Msg           string
	KeysAndValues []any
}

func (m *MockLogger) Debug(msg string, keysAndValues ...any) {
	m.DebugCalls = append(m.DebugCalls, logCall{Msg: msg, KeysAndValues: keysAndValues})
}

func (m *MockLogger) Info(msg string, keysAndValues ...any) {
	m.InfoCalls = append(m.InfoCalls, logCall{Msg: msg, KeysAndValues: keysAndValues})
}

func (m *MockLogger) Warn(msg string, keysAndValues ...any) {
	m.WarnCalls = append(m.WarnCalls, logCall{Msg: msg, KeysAndValues: keysAndValues})
}

func (m *MockLogger) Error(msg string, keysAndValues ...any) {
	m.ErrorCalls = append(m.ErrorCalls, logCall{Msg: msg, KeysAndValues: keysAndValues})
}

func TestLoggerAlias(t *testing.T) {
	t.Run("Logger alias implements logger.Handler", func(t *testing.T) {
		// Verify the type alias works correctly
		var l Logger = &MockLogger{}
		l.Debug("test")
		l.Info("test")
		l.Warn("test")
		l.Error("test")
	})
}

func TestWithDebug(t *testing.T) {
	t.Run("enables default logger when true", func(t *testing.T) {
		c := New("auth", "secret", WithDebug(true))

		// Debug should be enabled
		assert.True(t, c.Debug)

		// Should use default logger (not Nop)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)
	})

	t.Run("does not override custom logger", func(t *testing.T) {
		mockLogger := &MockLogger{}
		c := New("auth", "secret",
			WithLogger(mockLogger),
			WithDebug(true),
		)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)

		// Custom logger should have been used
		assert.NotEmpty(t, mockLogger.DebugCalls, "custom logger should be used")
	})
}

func TestWithLogger(t *testing.T) {
	t.Run("sets custom logger", func(t *testing.T) {
		mockLogger := &MockLogger{}
		c := New("auth", "secret", WithLogger(mockLogger))

		// Verify logger is set by making a request
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)

		// Should have logged debug and info messages
		assert.NotEmpty(t, mockLogger.DebugCalls, "expected debug logs")
		assert.NotEmpty(t, mockLogger.InfoCalls, "expected info logs")
	})

	t.Run("nil logger defaults to Nop", func(t *testing.T) {
		c := New("auth", "secret", WithLogger(nil))

		// Should not panic when making requests
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)
	})

	t.Run("default client uses Nop logger", func(t *testing.T) {
		c := New("auth", "secret")

		// Should not panic when making requests
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)
	})

	t.Run("works with logger subpackage directly", func(t *testing.T) {
		var buf bytes.Buffer
		l := logger.NewStd(&buf, logger.LevelDebug)

		c := New("auth", "secret", WithLogger(l))

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c.BaseURL = server.URL
		_, err := c.Get(t.Context(), "/test", nil)
		require.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "[DEBUG]")
		assert.Contains(t, output, "[INFO]")
	})
}

func TestLoggerIntegration(t *testing.T) {
	t.Run("logs request lifecycle", func(t *testing.T) {
		mockLogger := &MockLogger{}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{"result":"ok"}`,
			})
		}))
		defer server.Close()

		c := New("auth", "secret",
			WithBaseURL(server.URL),
			WithLogger(mockLogger),
		)

		_, err := c.Get(t.Context(), "/api/test", nil)
		require.NoError(t, err)

		// Verify debug log for sending request
		foundSendingRequest := false
		for _, call := range mockLogger.DebugCalls {
			if strings.Contains(call.Msg, "sending request") {
				foundSendingRequest = true
				break
			}
		}
		assert.True(t, foundSendingRequest, "expected 'sending request' debug log")

		// Verify info log for successful completion
		foundCompleted := false
		for _, call := range mockLogger.InfoCalls {
			if strings.Contains(call.Msg, "completed successfully") {
				foundCompleted = true
				break
			}
		}
		assert.True(t, foundCompleted, "expected 'completed successfully' info log")
	})

	t.Run("logs HTTP errors", func(t *testing.T) {
		mockLogger := &MockLogger{}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal error"))
		}))
		defer server.Close()

		c := New("auth", "secret",
			WithBaseURL(server.URL),
			WithLogger(mockLogger),
			WithRetries(0), // Disable retries for faster test
		)

		_, err := c.Get(t.Context(), "/api/test", nil)
		require.Error(t, err)

		// Verify error was logged
		assert.NotEmpty(t, mockLogger.ErrorCalls, "expected error logs for HTTP 500")
	})

	t.Run("logs retry attempts", func(t *testing.T) {
		mockLogger := &MockLogger{}
		attempts := 0

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attempts++
			if attempts < 2 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"code":    200,
				"message": "success",
				"data":    `{}`,
			})
		}))
		defer server.Close()

		c := New("auth", "secret",
			WithBaseURL(server.URL),
			WithLogger(mockLogger),
			WithRetries(2),
			WithRetryWait(1, 1), // Minimal wait for faster test
		)

		_, err := c.Get(t.Context(), "/api/test", nil)
		require.NoError(t, err)

		// Verify retry warning was logged
		foundRetry := false
		for _, call := range mockLogger.WarnCalls {
			if strings.Contains(call.Msg, "retry") {
				foundRetry = true
				break
			}
		}
		assert.True(t, foundRetry, "expected retry warning log")
	})
}

func TestDebugModeMissingLogs(t *testing.T) {
	t.Run("Debug mode suppresses Error logs", func(t *testing.T) {
		mockLogger := &MockLogger{}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		c := New("auth", "secret",
			WithBaseURL(server.URL),
			WithLogger(mockLogger),
			WithDebug(true), // Enable debug mode
			WithRetries(0),
		)

		_, err := c.Get(t.Context(), "/test", nil)
		require.Error(t, err)

		// We expect Error logs even in Debug mode
		assert.NotEmpty(t, mockLogger.ErrorCalls, "Expected Error logs to be present in Debug mode")

		// And we DO get Debug logs
		assert.NotEmpty(t, mockLogger.DebugCalls, "Expected Debug logs")
	})
}
