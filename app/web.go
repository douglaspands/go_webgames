package app

import (
	"github.com/gin-gonic/gin"
)

func CreateApp() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	controller := NewController()
	r.GET("/", controller.GetIndex)
	r.POST("/gameplay", controller.GameplayRedirect)
	r.GET("/gameplay/:console/:game", controller.Gameplay)
	r.GET("/roms", controller.RomList)
	r.Handle("GET", "/roms/download/:path", controller.RomDownload)
	r.Handle("HEAD", "/roms/download/:path", controller.RomDownload)
	r.Handle("GET", "/bios/download/:path", controller.BiosDownload)
	r.Handle("HEAD", "/bios/download/:path", controller.BiosDownload)

	return r
}
