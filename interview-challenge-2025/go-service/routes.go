package main

import (
	"github.com/gin-gonic/gin"
	"sort"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/packages", GetPackages)
	r.GET("/packages/:tracking_id", GetPackageByID)
	r.GET("/carriers", GetCarriers)
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

	c.JSON(200, packages)
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