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

// Handler defines the interface for structured logging in the SDK.
//
// This interface is compatible with most popular logging libraries:
//   - log/slog: Use slog.Logger directly (Go 1.21+)
//   - zerolog: Wrap with a simple adapter
//   - zap: Use zap.SugaredLogger
//   - logrus: Use logrus.Logger directly
//
// Example with log/slog:
//
//	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
//	client.New("auth", "secret", client.WithLogger(&SlogAdapter{slogger}))
//
// Example with custom logger:
//
//	type MyLogger struct{}
//	func (l *MyLogger) Debug(msg string, keysAndValues ...any) { /* ... */ }
//	func (l *MyLogger) Info(msg string, keysAndValues ...any)  { /* ... */ }
//	func (l *MyLogger) Warn(msg string, keysAndValues ...any)  { /* ... */ }
//	func (l *MyLogger) Error(msg string, keysAndValues ...any) { /* ... */ }
type Handler interface {
	// Debug logs a message at debug level with optional key-value pairs.
	Debug(msg string, keysAndValues ...any)
	// Info logs a message at info level with optional key-value pairs.
	Info(msg string, keysAndValues ...any)
	// Warn logs a message at warn level with optional key-value pairs.
	Warn(msg string, keysAndValues ...any)
	// Error logs a message at error level with optional key-value pairs.
	Error(msg string, keysAndValues ...any)
}

// Level represents the logging level.
type Level int

const (
	// LevelDebug enables all log messages.
	LevelDebug Level = iota
	// LevelInfo enables info, warn, and error messages.
	LevelInfo
	// LevelWarn enables warn and error messages.
	LevelWarn
	// LevelError enables only error messages.
	LevelError
	// LevelNone disables all logging.
	LevelNone
)

// String returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelNone:
		return "NONE"
	default:
		return "UNKNOWN"
	}
}
