package cwa_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omegaatt36/go-cwa"
)

func TestNewClient(t *testing.T) {
	client := cwa.NewClient("test-key")
	if client == nil {
		t.Fatal("client is nil")
	}
}

func TestGet36hForecast_Mock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":"true","result":{"resource_id":"F-C0032-001"},"records":{"datasetDescription":"36h forecast","location":[{"locationName":"臺北市"}]}}`))
	}))
	defer server.Close()

	client := cwa.NewClient("test-key", cwa.WithBaseURL(server.URL))
	resp, err := client.Get36hForecast(context.Background(), &cwa.Forecast36hParams{
		LocationNames: []cwa.County{cwa.TaipeiCity},
	})
	if err != nil {
		t.Fatalf("Get36hForecast failed: %v", err)
	}
	if !resp.IsSuccess() {
		t.Errorf("expected success true, got %s", resp.Success)
	}
	if loc := resp.Records.FirstLocation(); loc == nil || loc.LocationName != "臺北市" {
		t.Errorf("expected Taipei City location, got %v", loc)
	}
}

func TestGetTownshipForecast_Mock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":"true"}`))
	}))
	defer server.Close()

	client := cwa.NewClient("test-key", cwa.WithBaseURL(server.URL))
	resp, err := client.GetTownshipForecast(context.Background(), cwa.TownshipForecastParams{
		County: cwa.TaipeiCity,
		Period: cwa.ThreeDays,
	})
	if err != nil {
		t.Fatalf("GetTownshipForecast failed: %v", err)
	}
	if !resp.IsSuccess() {
		t.Errorf("expected success true, got %s", resp.Success)
	}
}
