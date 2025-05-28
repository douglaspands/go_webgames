package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func CreateApp() *gin.Engine {
	mode := os.Getenv("GIN_MODE")
	if len(mode) == 0 {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	controller := NewController()
	r.GET("/", controller.GetIndex)
	r.POST("/gameplay", controller.GameplayRedirect)
	r.GET("/gameplay/:console/:game", controller.Gameplay)
	r.GET("/roms", controller.ListGames)
	r.Handle("HEAD", "/download/:path", controller.Download)
	r.Handle("GET", "/download/:path", controller.Download)

	return r
}
