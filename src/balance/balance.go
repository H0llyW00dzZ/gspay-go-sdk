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

package balance

import (
	"context"
	"fmt"

	"github.com/H0llyW00dzZ/gspay-go-sdk/src/client"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/constants"
	"github.com/H0llyW00dzZ/gspay-go-sdk/src/i18n"
)

// Response represents the response from querying operator balance.
type Response struct {
	// Balance is the operator's IDR balance.
	Balance float64 `json:"balance"`
	// UsdtBalance is the operator's USDT balance.
	UsdtBalance float64 `json:"usdt_balance"`
}

// Service handles balance operations.
type Service struct{ client *client.Client }

// NewService creates a new balance service.
func NewService(c *client.Client) *Service { return &Service{client: c} }

// Get queries the operator's available settlement balance.
func (s *Service) Get(ctx context.Context) (*Response, error) {
	s.client.Logger().Debug(s.client.I18n(i18n.LogQueryingBalance))

	endpoint := fmt.Sprintf(constants.GetEndpoint(constants.EndpointBalance), s.client.AuthKey)
	resp, err := s.client.Get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	result, err := client.ParseData[Response](resp.Data, s.client.Language)
	if err != nil {
		return nil, err
	}

	s.client.Logger().Info(s.client.I18n(i18n.LogBalanceRetrieved),
		"idr_balance", result.Balance,
		"usdt_balance", result.UsdtBalance,
	)

	return result, nil
}
