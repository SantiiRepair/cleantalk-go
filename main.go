package main

import (
	"github.com/SantiiRepair/cleantalk-go/cleantalk"
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(cors.New(config))

	cleantalk.Handler(r)

	r.Run(":8080")
}
