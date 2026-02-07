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

// Package main demonstrates how to use the GSPAY Go SDK with proxy
// configurations via [client.WithHTTPClient].
//
// This example covers three common proxy patterns:
//   - HTTP/HTTPS proxy using a proxy URL
//   - Environment-based proxy using HTTP_PROXY/HTTPS_PROXY variables
//   - SOCKS5 proxy using [golang.org/x/net/proxy]
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/balance"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client"
)

func main() {
	// Get credentials from environment variables
	authKey := os.Getenv("GSPAY_AUTH_KEY")
	secretKey := os.Getenv("GSPAY_SECRET_KEY")

	if authKey == "" || secretKey == "" {
		log.Fatal("Please set GSPAY_AUTH_KEY and GSPAY_SECRET_KEY environment variables")
	}

	// Choose a proxy mode based on the PROXY_MODE environment variable.
	// Supported values: "http", "env", "socks5"
	proxyMode := os.Getenv("PROXY_MODE")
	if proxyMode == "" {
		proxyMode = "http"
	}

	var httpClient *http.Client
	switch proxyMode {
	case "http":
		httpClient = httpProxyClient()
	case "env":
		httpClient = envProxyClient()
	case "socks5":
		httpClient = socks5ProxyClient()
	default:
		log.Fatalf("Unknown PROXY_MODE: %s (supported: http, env, socks5)", proxyMode)
	}

	// Create the GSPAY client with the proxied HTTP client
	c := client.New(
		authKey,
		secretKey,
		client.WithHTTPClient(httpClient),
		client.WithTimeout(60*time.Second),
	)

	// Verify connectivity by checking balance
	ctx := context.Background()
	balanceSvc := balance.NewService(c)

	fmt.Printf("=== Checking Balance (proxy mode: %s) ===\n", proxyMode)
	resp, err := balanceSvc.Get(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("IDR Balance: %.2f\n", resp.Balance)
	fmt.Printf("USDT Balance: %.2f\n", resp.UsdtBalance)
}

// httpProxyClient creates an HTTP client that routes requests through
// an HTTP/HTTPS proxy specified by the PROXY_URL environment variable.
//
// Example:
//
//	PROXY_URL=http://proxy.example.com:8080
//	PROXY_URL=http://user:password@proxy.example.com:8080
func httpProxyClient() *http.Client {
	proxyAddr := os.Getenv("PROXY_URL")
	if proxyAddr == "" {
		log.Fatal("Please set PROXY_URL environment variable (e.g., http://proxy.example.com:8080)")
	}

	proxyURL, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatalf("Invalid PROXY_URL: %v", err)
	}

	fmt.Printf("Using HTTP proxy: %s\n", proxyURL.Host)

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
}

// envProxyClient creates an HTTP client that reads proxy settings
// from standard environment variables (HTTP_PROXY, HTTPS_PROXY, NO_PROXY).
//
// This is useful when the proxy is managed at the infrastructure level
// (e.g., Kubernetes, Docker, or CI/CD pipelines).
//
// Example:
//
//	HTTP_PROXY=http://proxy.example.com:8080
//	HTTPS_PROXY=http://proxy.example.com:8080
//	NO_PROXY=localhost,127.0.0.1
func envProxyClient() *http.Client {
	fmt.Println("Using proxy from environment variables (HTTP_PROXY/HTTPS_PROXY)")

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
}

// socks5ProxyClient creates an HTTP client that routes requests through
// a SOCKS5 proxy specified by the SOCKS5_PROXY environment variable.
//
// This uses Go's standard library [net.Dialer] with the SOCKS5 proxy address.
// For authenticated SOCKS5 proxies, use the golang.org/x/net/proxy package instead.
//
// Example:
//
//	SOCKS5_PROXY=127.0.0.1:1080
func socks5ProxyClient() *http.Client {
	socks5Addr := os.Getenv("SOCKS5_PROXY")
	if socks5Addr == "" {
		log.Fatal("Please set SOCKS5_PROXY environment variable (e.g., 127.0.0.1:1080)")
	}

	fmt.Printf("Using SOCKS5 proxy: %s\n", socks5Addr)

	// Use the socks5:// scheme with http.ProxyURL for SOCKS5 support
	// (available since Go 1.24).
	proxyURL, err := url.Parse("socks5://" + socks5Addr)
	if err != nil {
		log.Fatalf("Invalid SOCKS5_PROXY address: %v", err)
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
}
