package main

import "time"

// Package represents a tracked package
type Package struct {
	TrackingID   string    `json:"tracking_id"`
	Status       string    `json:"status"`
	Carrier      string    `json:"carrier"`
	Eta          time.Time `json:"eta"`
	LastUpdated  time.Time `json:"last_updated"`
	CurrentCity  string    `json:"current_city"`
}

// Carrier represents a shipping carrier
type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Location represents city/location metadata (for stretch goal)
type Location struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
} 