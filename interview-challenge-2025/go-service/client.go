package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const mockAPIBase = "http://localhost:8080"

var fetchAllPackages = realFetchAllPackages
var fetchPackageByID = realFetchPackageByID
var fetchCarriers = realFetchCarriers

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

	var carriers []Carrier
	if err := json.NewDecoder(resp.Body).Decode(&carriers); err != nil {
		return nil, err
	}
	return carriers, nil
} 