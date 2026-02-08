package client

import (
	"net"
	"net/http"
	"time"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client/logger"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/constants"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/i18n"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/internal/signature"
)

// Client is the GSPAY2 API client.
type Client struct {
	// AuthKey is the operator authentication key (used in URL path).
	AuthKey string
	// SecretKey is the operator secret key (used for signature generation).
	SecretKey string
	// BaseURL is the API base URL.
	BaseURL string
	// HTTPClient is the underlying HTTP client.
	// See [WithHTTPClient] for configuration.
	HTTPClient *http.Client
	// Timeout is the request timeout duration.
	Timeout time.Duration
	// Retries is the number of retry attempts for transient failures.
	Retries int
	// RetryWaitMin is the minimum wait time between retries.
	RetryWaitMin time.Duration
	// RetryWaitMax is the maximum wait time between retries.
	RetryWaitMax time.Duration
	// CallbackIPWhitelist contains allowed IP addresses/CIDR ranges for callbacks.
	// If empty, IP validation is skipped.
	CallbackIPWhitelist []string
	// parsedIPNets contains parsed CIDR networks for efficient IP checking.
	parsedIPNets []*net.IPNet
	// Debug controls whether sensitive data is sanitized in log output.
	//
	// If Debug is true, raw values (auth keys, account numbers, account names) are shown in logs,
	// and a default logger is used when no custom logger is set.
	// If Debug is false (default), sensitive data is automatically redacted for safe logging.
	Debug bool
	// parsedIPs contains parsed individual IP addresses.
	parsedIPs []net.IP
	// Language is the language for SDK error and log messages.
	// Default is [i18n.English]. See [WithLanguage] for configuration.
	Language i18n.Language
	// logger is the structured logger for the client.
	// Default is [logger.Nop] (no logging). See [WithLogger] for configuration.
	logger Logger
	// digest is the hash function for signature generation.
	// Default is nil (uses [crypto/md5]). See [WithDigest] for configuration.
	digest signature.Digest
	// qrOpts holds options for QR code generation, applied during initialization.
	// See [WithQRCodeOptions] for configuration.
	qrOpts []QROption
	// qrCfg holds the resolved QR code configuration.
	qrCfg *qrConfig
}

// New creates a new GSPAY2 API client.
//
// Parameters:
//   - authKey: Operator authentication key (used in URL path)
//   - secretKey: Operator secret key (used for signature generation)
//   - opts: Optional configuration options (see [Option])
func New(authKey, secretKey string, opts ...Option) *Client {
	c := &Client{
		AuthKey:      authKey,
		SecretKey:    secretKey,
		BaseURL:      constants.DefaultBaseURL,
		Timeout:      time.Duration(constants.DefaultTimeout) * time.Second,
		Retries:      constants.DefaultRetries,
		RetryWaitMin: time.Duration(constants.DefaultRetryWaitMin) * time.Millisecond,
		RetryWaitMax: time.Duration(constants.DefaultRetryWaitMax) * time.Millisecond,
		Language:     i18n.English,
		logger:       logger.Nop{},
		digest:       nil, // nil by default; explicit assignment for clarity (uses MD5)
		qrOpts:       nil, // nil by default; uses QR defaults (256px, Medium recovery)
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{
			Timeout: c.Timeout,
		}
	}

	// Initialize QR code config with configured options.
	c.qrCfg = qrDefaults()
	for _, opt := range c.qrOpts {
		opt(c.qrCfg)
	}

	return c
}
