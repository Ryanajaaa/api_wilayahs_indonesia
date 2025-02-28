package main

import (
	"github.com/Ryanajaa/go-restapi-gin/controllers"
	"github.com/Ryanajaa/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// Routes
	r.GET("/api/provinces/fetch", controllers.FetchProvinces) // Ambil dari API dan simpan ke DB
	r.GET("/api/provinces", controllers.GetProvinces)         // Ambil dari DB

	r.Run(":8080")
}
