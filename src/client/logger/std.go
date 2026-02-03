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
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Std is a simple logger that writes to the standard library log package.
// It formats messages with key-value pairs in a human-readable format.
type Std struct {
	logger *log.Logger
	level  Level
}

// NewStd creates a new [Std] logger with the specified output and log level.
//
// Example:
//
//	// Log to stdout at debug level
//	logger := logger.NewStd(os.Stdout, logger.LevelDebug)
//	c := client.New("auth", "secret", client.WithLogger(logger))
//
//	// Log to a file at info level
//	f, _ := os.Create("sdk.log")
//	logger := logger.NewStd(f, logger.LevelInfo)
func NewStd(w io.Writer, level Level) *Std {
	return &Std{
		logger: log.New(w, "", log.LstdFlags),
		level:  level,
	}
}

// Default returns a [Std] logger that writes to [os.Stderr] at [LevelDebug].
func Default() *Std {
	return NewStd(os.Stderr, LevelDebug)
}

// Debug implements [Handler.Debug].
func (l *Std) Debug(msg string, keysAndValues ...any) {
	if l.level <= LevelDebug {
		l.log("DEBUG", msg, keysAndValues...)
	}
}

// Info implements [Handler.Info].
func (l *Std) Info(msg string, keysAndValues ...any) {
	if l.level <= LevelInfo {
		l.log("INFO", msg, keysAndValues...)
	}
}

// Warn implements [Handler.Warn].
func (l *Std) Warn(msg string, keysAndValues ...any) {
	if l.level <= LevelWarn {
		l.log("WARN", msg, keysAndValues...)
	}
}

// Error implements [Handler.Error].
func (l *Std) Error(msg string, keysAndValues ...any) {
	if l.level <= LevelError {
		l.log("ERROR", msg, keysAndValues...)
	}
}

// log formats and writes a log message with key-value pairs.
func (l *Std) log(level, msg string, keysAndValues ...any) {
	if len(keysAndValues) == 0 {
		l.logger.Printf("[%s] %s", level, msg)
		return
	}

	// Format key-value pairs
	kvStr := FormatKeyValues(keysAndValues...)
	l.logger.Printf("[%s] %s %s", level, msg, kvStr)
}

// FormatKeyValues formats key-value pairs as "key1=value1 key2=value2".
func FormatKeyValues(keysAndValues ...any) string {
	if len(keysAndValues) == 0 {
		return ""
	}

	// Handle odd number of arguments
	if len(keysAndValues)%2 != 0 {
		keysAndValues = append(keysAndValues, "(MISSING)")
	}

	var result strings.Builder
	for i := 0; i < len(keysAndValues); i += 2 {
		key := fmt.Sprintf("%v", keysAndValues[i])
		value := keysAndValues[i+1]

		if i > 0 {
			result.WriteString(" ")
		}
		result.WriteString(fmt.Sprintf("%s=%v", key, value))
	}
	return result.String()
}

// Level returns the current log level.
func (l *Std) Level() Level {
	return l.level
}
