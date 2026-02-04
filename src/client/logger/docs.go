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

// Package logger provides structured logging interfaces and implementations
// for the GSPAY2 SDK client.
//
// The package defines a [Handler] interface that is compatible with popular
// logging libraries including log/slog (Go 1.21+), zap, zerolog, and logrus.
//
// # Basic Usage
//
// The SDK defaults to [Nop] which discards all log messages.
// To enable logging, use client.WithLogger:
//
//	import "github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
//
//	// Use the default stderr logger
//	c := client.New("auth", "secret",
//	    client.WithLogger(logger.Default()),
//	)
//
//	// Use a custom log level
//	c := client.New("auth", "secret",
//	    client.WithLogger(logger.NewStd(os.Stdout, logger.LevelInfo)),
//	)
//
// # Custom Logger Implementation
//
// Implement the [Handler] interface to use your preferred logging library:
//
//	type SlogAdapter struct {
//	    logger *slog.Logger
//	}
//
//	func (a *SlogAdapter) Debug(msg string, keysAndValues ...any) {
//	    a.logger.Debug(msg, keysAndValues...)
//	}
//	// ... implement Info, Warn, Error
//
// # Log Levels
//
// The [Std] logger supports the following levels:
//   - [LevelDebug]: All messages (most verbose)
//   - [LevelInfo]: Info, Warn, Error messages
//   - [LevelWarn]: Warn, Error messages
//   - [LevelError]: Error messages only
//   - [LevelNone]: No messages (silent)
package logger
