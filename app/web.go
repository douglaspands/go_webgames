package app

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// A função createApp é responsável por criar um aplicativo web.
func CreateApp() *gin.Engine {
	r := gin.Default()

	templatesDir := filepath.Dir("./templates")
	staticDir := filepath.Dir("./static")

	r.LoadHTMLGlob(templatesDir + "/*.html")
	r.Static("/static", staticDir)
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
