package controllers

import (
	"time"

	"arna-cricket/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	controller := controllers.NewController()
	r := gin.Default()
	r.UseRawPath = true
	r.UnescapePathValues = false

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS, HEAD, PATCH"},
		AllowHeaders:     []string{"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cric := r.Group("api/v1/cricket")
	{
		cric.GET("/scorecard", controller.CricInfoController.GetScorecard)
	}

	return r
}
