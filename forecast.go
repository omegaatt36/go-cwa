package cwa

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const cwaTimeFormat = "2006-01-02T15:04:05"

func addTimeParams(v url.Values, startTime, endTime, timeFrom, timeTo *time.Time) {
	if startTime != nil {
		v.Set("startTime", startTime.Format(cwaTimeFormat))
	}
	if endTime != nil {
		v.Set("endTime", endTime.Format(cwaTimeFormat))
	}
	if timeFrom != nil {
		v.Set("timeFrom", timeFrom.Format(cwaTimeFormat))
	}
	if timeTo != nil {
		v.Set("timeTo", timeTo.Format(cwaTimeFormat))
	}
}

// Forecast36hParams defines parameters for the 36-hour forecast API.
type Forecast36hParams struct {
	LocationNames []County
	ElementNames  []string
	StartTime     *time.Time
	EndTime       *time.Time
	TimeFrom      *time.Time
	TimeTo        *time.Time
}

// Get36hForecast returns the 36-hour weather forecast.
func (c *Client) Get36hForecast(ctx context.Context, params *Forecast36hParams) (*Response[Forecast36hRecords], error) {
	v := url.Values{}
	v.Set("Authorization", c.apiKey)

	if params != nil {
		if len(params.LocationNames) > 0 {
			locs := make([]string, len(params.LocationNames))
			for i, l := range params.LocationNames {
				locs[i] = string(l)
			}
			v.Set("locationName", strings.Join(locs, ","))
		}
		if len(params.ElementNames) > 0 {
			v.Set("elementName", strings.Join(params.ElementNames, ","))
		}
		addTimeParams(v, params.StartTime, params.EndTime, params.TimeFrom, params.TimeTo)
	}

	u := fmt.Sprintf("%s/v1/rest/datastore/F-C0032-001?%s", c.baseURL, v.Encode())

	var resp Response[Forecast36hRecords]
	if err := c.doRequest(ctx, u, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// TownshipForecastParams defines parameters for the township forecast API.
type TownshipForecastParams struct {
	County        County
	Period        Period
	LocationNames []string
	StartTime     *time.Time
	EndTime       *time.Time
	TimeFrom      *time.Time
	TimeTo        *time.Time
}

// GetTownshipForecast returns the township forecast for the given county and period.
func (c *Client) GetTownshipForecast(ctx context.Context, params *TownshipForecastParams) (*Response[map[string]any], error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	countyMap, ok := townshipDatasetIDs[params.County]
	if !ok {
		return nil, fmt.Errorf("unsupported county: %s", params.County)
	}
	datasetID, ok := countyMap[params.Period]
	if !ok {
		return nil, fmt.Errorf("unsupported period: %s", params.Period)
	}

	v := url.Values{}
	v.Set("Authorization", c.apiKey)
	if len(params.LocationNames) > 0 {
		v.Set("locationName", strings.Join(params.LocationNames, ","))
	}
	addTimeParams(v, params.StartTime, params.EndTime, params.TimeFrom, params.TimeTo)

	u := fmt.Sprintf("%s/v1/rest/datastore/%s?%s", c.baseURL, datasetID, v.Encode())

	var resp Response[map[string]any]
	if err := c.doRequest(ctx, u, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) doRequest(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
