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
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
)

// Logger is an alias for [logger.Handler].
//
// This allows users to implement custom loggers without importing the logger subpackage.
//
// Example:
//
//	type MyLogger struct{}
//	func (l *MyLogger) Debug(msg string, keysAndValues ...any) { /* ... */ }
//	func (l *MyLogger) Info(msg string, keysAndValues ...any)  { /* ... */ }
//	func (l *MyLogger) Warn(msg string, keysAndValues ...any)  { /* ... */ }
//	func (l *MyLogger) Error(msg string, keysAndValues ...any) { /* ... */ }
//
//	c := client.New("auth", "secret", client.WithLogger(&MyLogger{}))
type Logger = logger.Handler

// Logger returns the configured logger instance.
//
// This allows services and other packages to use the same logger configuration
// as the client for consistent logging across the SDK.
func (c *Client) Logger() Logger {
	return c.logger
}
