package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
} var fetchAllPackages = realFetchAllPackages
var fetchPackageByID = realFetchPackageByID
var fetchCarriers = realFetchCarriers