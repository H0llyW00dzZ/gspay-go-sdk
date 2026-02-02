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

// Package main demonstrates logging options with the GSPAY Go SDK.
//
// This example shows:
//   - Using WithDebug for simple debug logging
//   - Using WithLogger with the built-in Std logger
//   - Creating a custom logger adapter for slog
//   - Different log levels
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/balance"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
)

func main() {
	// Get credentials from environment variables
	authKey := os.Getenv("GSPAY_AUTH_KEY")
	secretKey := os.Getenv("GSPAY_SECRET_KEY")

	if authKey == "" || secretKey == "" {
		fmt.Println("Please set GSPAY_AUTH_KEY and GSPAY_SECRET_KEY environment variables")
		fmt.Println("Running with demo mode...")
		authKey = "demo-auth-key"
		secretKey = "demo-secret-key"
	}

	ctx := context.Background()

	// Example 1: Simple debug mode with WithDebug
	fmt.Println("=== Example 1: WithDebug (Debug-only output) ===")
	example1DebugMode(ctx, authKey, secretKey)

	fmt.Println()

	// Example 2: Built-in Std logger with different levels
	fmt.Println("=== Example 2: Std Logger (All levels) ===")
	example2StdLogger(ctx, authKey, secretKey)

	fmt.Println()

	// Example 3: Custom slog adapter
	fmt.Println("=== Example 3: Custom slog Adapter (JSON output) ===")
	example3SlogAdapter(ctx, authKey, secretKey)

	fmt.Println()

	// Example 4: File logging
	fmt.Println("=== Example 4: File Logging ===")
	example4FileLogging(ctx, authKey, secretKey)
}

// example1DebugMode demonstrates the simple WithDebug option.
// Only DEBUG level messages are shown (no INFO/WARN/ERROR).
func example1DebugMode(ctx context.Context, authKey, secretKey string) {
	c := client.New(
		authKey,
		secretKey,
		client.WithDebug(true), // Simple debug mode
		client.WithTimeout(10*time.Second),
	)

	balanceSvc := balance.NewService(c)
	_, err := balanceSvc.Get(ctx)
	if err != nil {
		fmt.Printf("Expected error (demo mode): %v\n", err)
	}
}

// example2StdLogger demonstrates the built-in Std logger with custom log levels.
func example2StdLogger(ctx context.Context, authKey, secretKey string) {
	// Create a logger that writes to stdout at Info level
	// This will show INFO, WARN, and ERROR but not DEBUG
	stdLogger := logger.NewStd(os.Stdout, logger.LevelInfo)

	c := client.New(
		authKey,
		secretKey,
		client.WithLogger(stdLogger),
		client.WithTimeout(10*time.Second),
	)

	balanceSvc := balance.NewService(c)
	_, err := balanceSvc.Get(ctx)
	if err != nil {
		fmt.Printf("Expected error (demo mode): %v\n", err)
	}
}

// example3SlogAdapter demonstrates using Go's log/slog with a custom adapter.
func example3SlogAdapter(ctx context.Context, authKey, secretKey string) {
	// Create a slog logger with JSON output
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Wrap it with our adapter
	adapter := &SlogAdapter{Logger: slogger}

	c := client.New(
		authKey,
		secretKey,
		client.WithLogger(adapter),
		client.WithTimeout(10*time.Second),
	)

	balanceSvc := balance.NewService(c)
	_, err := balanceSvc.Get(ctx)
	if err != nil {
		fmt.Printf("Expected error (demo mode): %v\n", err)
	}
}

// example4FileLogging demonstrates logging to a file.
func example4FileLogging(ctx context.Context, authKey, secretKey string) {
	// Create a log file in the current directory
	logFileName := "logging.example"
	f, err := os.Create(logFileName)
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Printf("Logging to: %s\n", logFileName)

	// Create a logger that writes to the file
	fileLogger := logger.NewStd(f, logger.LevelDebug)

	c := client.New(
		authKey,
		secretKey,
		client.WithLogger(fileLogger),
		client.WithTimeout(10*time.Second),
	)

	balanceSvc := balance.NewService(c)
	_, _ = balanceSvc.Get(ctx)

	// Read and display the log file contents
	f.Seek(0, 0)
	content := make([]byte, 1024)
	n, _ := f.Read(content)
	fmt.Printf("Log file contents:\n%s\n", string(content[:n]))
}

// SlogAdapter adapts Go's log/slog to the logger.Handler interface.
//
// This demonstrates how to integrate the SDK with popular logging libraries.
type SlogAdapter struct {
	Logger *slog.Logger
}

// Debug implements logger.Handler.
func (a *SlogAdapter) Debug(msg string, keysAndValues ...any) {
	a.Logger.Debug(msg, keysAndValues...)
}

// Info implements logger.Handler.
func (a *SlogAdapter) Info(msg string, keysAndValues ...any) {
	a.Logger.Info(msg, keysAndValues...)
}

// Warn implements logger.Handler.
func (a *SlogAdapter) Warn(msg string, keysAndValues ...any) {
	a.Logger.Warn(msg, keysAndValues...)
}

// Error implements logger.Handler.
func (a *SlogAdapter) Error(msg string, keysAndValues ...any) {
	a.Logger.Error(msg, keysAndValues...)
}
