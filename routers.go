package main

import "github.com/gin-gonic/gin"

// SetupRouter : init routes
func SetupRouter() *gin.Engine {

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	g.Static("/web", "./web")
	g.Static("/assets", "./assets")
	g.Static("/node_modules", "./node_modules")
	g.StaticFile("/favicon.ico", "./assets/favicon.ico")
	g.LoadHTMLGlob("templates/*.html")

	// g := gin.New()

	// Logging middleware
	g.Use(gin.Logger())
	// Recovery middleware
	g.Use(gin.Recovery())

	/*
		g.Use(static.Serve("/web", static.LocalFile("/web", false)))
		v1 := router.Group("api/v1")
		{
			v1.GET("/instructions", GetInstructions)
		}
	*/

	// g.Use(permissionHandler)

	return g
}
