package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mustParseTime(s string) (t time.Time) {
	t, _ = time.Parse(time.RFC3339, s)
	return
}

func TestGetPackages_Basic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Mock fetchAllPackages
	fetchAllPackagesFunc := fetchAllPackages
	fetchAllPackages = func() ([]Package, error) {
		return []Package{
			{
				TrackingID:  "PKG1",
				Status:      "Delivered",
				Carrier:     "UPS",
				Eta:         time.Now(),
				LastUpdated: time.Now(),
				CurrentCity: "New York",
			},
		}, nil
	}
	defer func() { fetchAllPackages = fetchAllPackagesFunc }()

	r := gin.Default()
	r.GET("/packages", GetPackages)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/packages", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "PKG1")
}

func TestGetPackageRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Mock fetchPackageByID and fetchLocation
	fetchPackageByIDFunc := fetchPackageByID
	fetchLocationFunc := fetchLocation
	fetchPackageByID = func(id string) (*Package, error) {
		return &Package{
			TrackingID:  "PKG1",
			Status:      "Delivered",
			Carrier:     "UPS",
			Eta:         mustParseTime("2025-06-15T18:00:00Z"),
			LastUpdated: mustParseTime("2025-06-12T09:30:00Z"),
			CurrentCity: "New York",
		}, nil
	}
	fetchLocation = func(city string) (*Location, error) {
		cityData := map[string]Location{
			"Chicago":      {City: "Chicago", Latitude: 41.8781, Longitude: -87.6298},
			"Philadelphia": {City: "Philadelphia", Latitude: 39.9526, Longitude: -75.1652},
			"New York":     {City: "New York", Latitude: 40.7128, Longitude: -74.0060},
		}
		loc, ok := cityData[city]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		return &loc, nil
	}
	defer func() {
		fetchPackageByID = fetchPackageByIDFunc
		fetchLocation = fetchLocationFunc
	}()

	r := gin.Default()
	r.GET("/packages/:tracking_id/route", GetPackageRoute)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/packages/PKG1/route", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Chicago")
	assert.Contains(t, w.Body.String(), "New York")
	assert.Contains(t, w.Body.String(), "timestamp")
}

func TestGetPackageByID_Basic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	fetchPackageByIDFunc := fetchPackageByID
	fetchPackageByID = func(id string) (*Package, error) {
		if id == "PKG1" {
			return &Package{
				TrackingID:  "PKG1",
				Status:      "Delivered",
				Carrier:     "UPS",
				Eta:         time.Now(),
				LastUpdated: time.Now(),
				CurrentCity: "New York",
			}, nil
		}
		return nil, fmt.Errorf("not found")
	}
	defer func() { fetchPackageByID = fetchPackageByIDFunc }()

	r := gin.Default()
	r.GET("/packages/:tracking_id", GetPackageByID)

	t.Run("found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/packages/PKG1", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "PKG1")
	})

	t.Run("not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/packages/UNKNOWN", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
}

func TestGetCarriers_Basic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	fetchCarriersFunc := fetchCarriers
	fetchCarriers = func() ([]Carrier, error) {
		return []Carrier{{ID: "UPS", Name: "United Parcel Service"}}, nil
	}
	defer func() { fetchCarriers = fetchCarriersFunc }()

	r := gin.Default()
	r.GET("/carriers", GetCarriers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/carriers", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "UPS")
}