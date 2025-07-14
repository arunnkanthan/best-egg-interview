package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const mockAPIBase = "http://localhost:8080"

var fetchAllPackages = realFetchAllPackages
var fetchPackageByID = realFetchPackageByID
var fetchCarriers = realFetchCarriers
var fetchLocation = realFetchLocation

func realFetchAllPackages() ([]Package, error) {
	resp, err := DoRequestWithRetry(func() (*http.Response, error) {
		return http.Get(fmt.Sprintf("%s/tracking", mockAPIBase))
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type packagesResponse struct {
		Packages []Package `json:"packages"`
	}
	var pr packagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return nil, err
	}
	return pr.Packages, nil
}

func realFetchPackageByID(id string) (*Package, error) {
	resp, err := DoRequestWithRetry(func() (*http.Response, error) {
		return http.Get(fmt.Sprintf("%s/tracking/%s", mockAPIBase, id))
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("not found")
	}

	var pkg Package
	if err := json.NewDecoder(resp.Body).Decode(&pkg); err != nil {
		return nil, err
	}
	return &pkg, nil
}

func realFetchCarriers() ([]Carrier, error) {
	resp, err := DoRequestWithRetry(func() (*http.Response, error) {
		return http.Get(fmt.Sprintf("%s/carriers", mockAPIBase))
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type carriersResponse struct {
		Carriers []Carrier `json:"carriers"`
	}
	var cr carriersResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return nil, err
	}
	return cr.Carriers, nil
}

func realFetchLocation(city string) (*Location, error) {
	url := fmt.Sprintf("%s/locations/%s", mockAPIBase, urlEncodeCity(city))
	resp, err := DoRequestWithRetry(func() (*http.Response, error) {
		return http.Get(url)
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var loc Location
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return nil, err
	}
	return &loc, nil
}

func urlEncodeCity(city string) string {
	// Replace spaces with %20 for URL encoding
	return strings.ReplaceAll(city, " ", "%20")
} 