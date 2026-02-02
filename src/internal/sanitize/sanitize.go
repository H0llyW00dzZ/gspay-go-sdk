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

package sanitize

import "strings"

// Endpoint redacts sensitive information like auth keys from endpoint URLs.
//
// This function handles the GSPAY2 API endpoint patterns:
//   - /v2/integrations/operator/{authkey}/...  (singular - e.g., balance)
//   - /v2/integrations/operators/{authkey}/... (plural - e.g., USDT)
//
// The auth key in position 4 is replaced with "[REDACTED]".
//
// Example:
//
//	sanitize.Endpoint("/v2/integrations/operators/secret123/idr/payment")
//	// Returns: "/v2/integrations/operators/[REDACTED]/idr/payment"
func Endpoint(endpoint string) string {
	// Path structure after split:
	// parts[0] = "" (empty, from leading slash)
	// parts[1] = "v2"
	// parts[2] = "integrations"
	// parts[3] = "operator" or "operators"
	// parts[4] = authkey (to be redacted)
	// parts[5+] = remaining path segments
	parts := strings.Split(endpoint, "/")
	if len(parts) >= 5 && parts[1] == "v2" && parts[2] == "integrations" && len(parts[4]) > 0 {
		if parts[3] == "operator" || parts[3] == "operators" {
			parts[4] = "[REDACTED]"
			return strings.Join(parts, "/")
		}
	}
	return endpoint
}
