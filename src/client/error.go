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

import "github.com/H0llyW00dzZ/gspay-go-sdk/src/errors"

// Error creates a localized error wrapping the provided sentinel error.
// This is a convenience method that uses the client's language setting.
func (c *Client) Error(sentinel error, args ...any) error {
	return errors.New(c.Language, sentinel, args...)
}
