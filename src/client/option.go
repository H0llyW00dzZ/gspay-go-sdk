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
	"net/http"
	"strings"
	"time"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/i18n"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/internal/signature"
)

// Option is a functional option for configuring the Client.
type Option func(*Client)

// WithBaseURL sets a custom base URL for the API.
//
// Trailing slashes are automatically trimmed from the URL.
// Default is "https://api.thegspay.com".
//
// Example:
//
//	c := client.New("auth", "secret", client.WithBaseURL("https://sandbox.api.com"))
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.BaseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithHTTPClient sets a custom HTTP client.
//
// Use this to configure custom transport settings, proxies, or TLS configurations.
// If not set, a default http.Client is used.
//
// Example:
//
//	customClient := &http.Client{
//	    Transport: &http.Transport{
//	        MaxIdleConns: 100,
//	    },
//	}
//	c := client.New("auth", "secret", client.WithHTTPClient(customClient))
//
// For advanced use cases such as routing requests through a proxy:
//
//	proxyURL, _ := url.Parse("http://proxy.example.com:8080")
//	customClient := &http.Client{
//	    Transport: &http.Transport{
//	        Proxy: http.ProxyURL(proxyURL),
//	    },
//	}
//	c := client.New("auth", "secret", client.WithHTTPClient(customClient))
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithTimeout sets the request timeout.
//
// The timeout applies to each individual HTTP request.
// Minimum allowed timeout is 5 seconds; values below this are ignored.
// Default is 30 seconds.
//
// Example:
//
//	c := client.New("auth", "secret", client.WithTimeout(60*time.Second))
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if timeout >= 5*time.Second {
			c.Timeout = timeout
		}
	}
}

// WithRetries sets the number of retry attempts for transient failures.
//
// Retries are attempted for 5xx server errors, timeouts, and connection issues.
// Negative values are ignored. Set to 0 to disable retries.
// Default is 3 retries.
//
// Example:
//
//	c := client.New("auth", "secret", client.WithRetries(5))
func WithRetries(retries int) Option {
	return func(c *Client) {
		if retries >= 0 {
			c.Retries = retries
		}
	}
}

// WithDebug enables debug mode for the client.
//
// When enabled, sensitive data (auth keys, account numbers, account names) is shown
// unsanitized in log output, and a default logger is automatically used if no custom
// logger is set.
//
// Example:
//
//	// Enable debug logging (uses default stderr logger)
//	c := client.New("auth", "secret", client.WithDebug(true))
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.Debug = debug
		if debug {
			// Use default logger if none is set
			if _, isNop := c.logger.(logger.Nop); isNop {
				c.logger = logger.Default()
			}
		}
	}
}

// WithRetryWait sets the minimum and maximum wait times between retries.
//
// The actual wait time is calculated using exponential backoff with jitter,
// bounded between min and max values.
// Default is 500ms minimum, 2s maximum.
//
// Example:
//
//	c := client.New("auth", "secret", client.WithRetryWait(1*time.Second, 5*time.Second))
func WithRetryWait(min, max time.Duration) Option {
	return func(c *Client) {
		c.RetryWaitMin = min
		c.RetryWaitMax = max
	}
}

// WithCallbackIPWhitelist sets the allowed IP addresses or CIDR ranges for callback verification.
//
// Accepts individual IP addresses (e.g., "192.168.1.1") or CIDR notation (e.g., "192.168.1.0/24").
// If the whitelist is empty, IP validation is skipped during callback verification.
//
// Example:
//
//	client.New("auth", "secret", client.WithCallbackIPWhitelist(
//	    "192.168.1.1",
//	    "10.0.0.0/8",
//	    "2001:db8::/32",
//	))
func WithCallbackIPWhitelist(ips ...string) Option {
	return func(c *Client) {
		c.CallbackIPWhitelist = ips
		c.parseIPWhitelist()
	}
}

// WithLanguage sets the language for localized SDK messages.
// This affects error messages, log messages, and the output of
// [Client.I18n] and [Client.Error] methods.
//
// Default is [i18n.English]. Supported languages:
//   - [i18n.English] - English (default)
//   - [i18n.Indonesian] - Indonesian (Bahasa Indonesia)
//
// Example:
//
//	client.New("auth", "secret", client.WithLanguage(i18n.Indonesian))
func WithLanguage(lang i18n.Language) Option {
	return func(c *Client) {
		if lang.IsValid() {
			c.Language = lang
		}
	}
}

// WithLogger sets a custom logger for the client.
//
// If l is nil, a [logger.Nop] is used (no logging).
// For debug logging, use [logger.Default] or [logger.NewStd].
//
// Example:
//
//	// Enable debug logging to stderr
//	c := client.New("auth", "secret", client.WithLogger(logger.Default()))
//
//	// Custom log level
//	l := logger.NewStd(os.Stdout, logger.LevelInfo)
//	c := client.New("auth", "secret", client.WithLogger(l))
//
//	// Disable logging explicitly
//	c := client.New("auth", "secret", client.WithLogger(nil))
func WithLogger(l logger.Handler) Option {
	return func(c *Client) {
		if l == nil {
			c.logger = logger.Nop{}
		} else {
			c.logger = l
		}
	}
}

// WithDigest sets a custom hash function for signature generation.
//
// Note: The GSPAY2 API requires MD5 for signature verification.
// This option is provided for flexibility in case the API adds support
// for stronger algorithms, or for use with other compatible APIs.
// If digest is nil, MD5 is used (default behavior for GSPAY2 compatibility).
//
// Standard library hash functions can be used directly:
//   - [crypto/md5.New] (default, required by GSPAY2)
//   - [crypto/sha1.New]
//   - [crypto/sha256.New]
//   - [crypto/sha512.New]
//
// Example:
//
//	// Use SHA-256 for signature generation (if API supports it)
//	c := client.New("auth", "secret", client.WithDigest(sha256.New))
//
//	// Use SHA-512
//	c := client.New("auth", "secret", client.WithDigest(sha512.New))
func WithDigest(digest signature.Digest) Option {
	return func(c *Client) {
		c.digest = digest
	}
}
