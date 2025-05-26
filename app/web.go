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

	r.GET("/", getIndex)
	r.POST("/gameplay", gameplayRedirect)
	r.GET("/gameplay/:console/:game", gameplay)
	r.GET("/roms", romList)
	r.Handle("GET", "/roms/download/*path", romDownload)
	r.Handle("HEAD", "/roms/download/*path", romDownload)
	r.Handle("GET", "/bios/download/*path", biosDownload)
	r.Handle("HEAD", "/bios/download/*path", biosDownload)

	return r
}
