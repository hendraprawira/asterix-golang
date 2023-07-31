package router

import (
	"asterix-golang/app/controllers/asterix"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // for release mode, uncomment if need it
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Define a route handler for handling HTTP requests
	apiUri := r.Group("")
	asterixRoute := apiUri.Group("")
	{
		asterixRoute.GET("/geosocket", asterix.WebSocket)
	}

	return r
}
