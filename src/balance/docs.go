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

// Package balance provides balance query functionality for the GSPAY2 SDK.
//
// This package allows merchants to check their settlement balance on the
// GSPAY2 platform for managing funds and reconciliation.
//
// # Basic Usage
//
//	c := client.New("auth-key", "secret-key")
//	balanceSvc := balance.NewService(c)
//
//	resp, err := balanceSvc.Get(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Balance: %s\n", resp.Balance)
//
// # Response
//
// The [Response] struct contains:
//   - Balance: The current settlement balance as a string
//
// # Error Handling
//
// The Get method may return:
//   - API errors (invalid credentials, server errors)
//   - Network errors (timeout, connection issues)
//
// Use the SDK errors package GetAPIError function to handle API-specific errors.
package balance
