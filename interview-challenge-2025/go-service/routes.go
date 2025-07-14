package main

import (
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
	"time"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/packages", GetPackages)
	r.GET("/packages/:tracking_id", GetPackageByID)
	r.GET("/carriers", GetCarriers)
	r.GET("/packages/:tracking_id/route", GetPackageRoute) // route is a simulation of the package's route, bonus requirement
}

// Handler stubs
func GetPackages(c *gin.Context) {
	packages, err := fetchAllPackages()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Filtering by status
	status := c.Query("status")
	if status != "" {
		filtered := make([]Package, 0)
		for _, p := range packages {
			if p.Status == status {
				filtered = append(filtered, p)
			}
		}
		packages = filtered
	}

	// Sorting by eta or last_updated
	sortBy := c.Query("sort")
	if sortBy == "eta" {
		sort.Slice(packages, func(i, j int) bool {
			return packages[i].Eta.Before(packages[j].Eta)
		})
	} else if sortBy == "last_updated" {
		sort.Slice(packages, func(i, j int) bool {
			return packages[i].LastUpdated.Before(packages[j].LastUpdated)
		})
	}

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	start := (page - 1) * limit
	end := start + limit
	if start > len(packages) {
		start = len(packages)
	}
	if end > len(packages) {
		end = len(packages)
	}
	paginated := packages[start:end]

	c.JSON(200, gin.H{
		"packages": paginated,
		"page": page,
		"limit": limit,
		"total": len(packages),
	})
}

func GetPackageByID(c *gin.Context) {
	id := c.Param("tracking_id")
	pkg, err := fetchPackageByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Package not found"})
		return
	}
	c.JSON(200, pkg)
}

func GetCarriers(c *gin.Context) {
	carriers, err := fetchCarriers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, carriers)
}

func GetPackageRoute(c *gin.Context) {
	id := c.Param("tracking_id")
	pkg, err := fetchPackageByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Package not found"})
		return
	}

	// Simulate a route: pick 2-3 cities including the current city
	cities := []string{"Chicago", "Philadelphia", pkg.CurrentCity}
	startTime := pkg.LastUpdated.Add(-48 * time.Hour) // simulate 2 days before last update
	eta := pkg.Eta
	interval := eta.Sub(startTime) / time.Duration(len(cities))

	route := make([]gin.H, 0, len(cities))
	for i, city := range cities {
		loc, err := fetchLocation(city)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch location for " + city})
			return
		}
		timestamp := startTime.Add(time.Duration(i) * interval)
		route = append(route, gin.H{
			"city":      loc.City,
			"lat":       loc.Latitude,
			"lon":       loc.Longitude,
			"timestamp": timestamp.Format(time.RFC3339),
		})
	}

	c.JSON(200, gin.H{"route": route})
} 