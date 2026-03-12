package cwa_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func TestGet36hForecast_WithTimeParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Check if startTime is correctly encoded in the query
		startTime := r.URL.Query().Get("startTime")
		if startTime != "2024-03-12T00:00:00" {
			t.Errorf("expected startTime 2024-03-12T00:00:00, got %s", startTime)
		}

		w.Write([]byte(`{"success":"true","records":{"location":[]}}`))
	}))
	defer server.Close()

	client := cwa.NewClient("test-key", cwa.WithBaseURL(server.URL))
	st, _ := time.Parse("2006-01-02T15:04:05", "2024-03-12T00:00:00")
	_, err := client.Get36hForecast(context.Background(), &cwa.Forecast36hParams{
		StartTime: &st,
	})
	if err != nil {
		t.Fatalf("Get36hForecast failed: %v", err)
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
	resp, err := client.GetTownshipForecast(context.Background(), &cwa.TownshipForecastParams{
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

func TestGetTownshipForecast_WithTimeParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		q := r.URL.Query()
		if q.Get("startTime") != "2024-03-12T00:00:00" {
			t.Errorf("expected startTime 2024-03-12T00:00:00, got %s", q.Get("startTime"))
		}
		if q.Get("endTime") != "2024-03-13T00:00:00" {
			t.Errorf("expected endTime 2024-03-13T00:00:00, got %s", q.Get("endTime"))
		}
		if q.Get("timeFrom") != "2024-03-11T00:00:00" {
			t.Errorf("expected timeFrom 2024-03-11T00:00:00, got %s", q.Get("timeFrom"))
		}
		if q.Get("timeTo") != "2024-03-14T00:00:00" {
			t.Errorf("expected timeTo 2024-03-14T00:00:00, got %s", q.Get("timeTo"))
		}

		w.Write([]byte(`{"success":"true"}`))
	}))
	defer server.Close()

	client := cwa.NewClient("test-key", cwa.WithBaseURL(server.URL))
	st, _ := time.Parse("2006-01-02T15:04:05", "2024-03-12T00:00:00")
	et, _ := time.Parse("2006-01-02T15:04:05", "2024-03-13T00:00:00")
	tf, _ := time.Parse("2006-01-02T15:04:05", "2024-03-11T00:00:00")
	tt, _ := time.Parse("2006-01-02T15:04:05", "2024-03-14T00:00:00")

	_, err := client.GetTownshipForecast(context.Background(), &cwa.TownshipForecastParams{
		County:    cwa.TaipeiCity,
		Period:    cwa.ThreeDays,
		StartTime: &st,
		EndTime:   &et,
		TimeFrom:  &tf,
		TimeTo:    &tt,
	})
	if err != nil {
		t.Fatalf("GetTownshipForecast failed: %v", err)
	}
}
