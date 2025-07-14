package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	// Swagger docs at /docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterRoutes(r)

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
} 