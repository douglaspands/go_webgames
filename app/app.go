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

	repository := NewRepository()
	service := NewService(repository)
	controller := NewController(service)

	r.Handle("GET", "/", controller.GetIndex)
	r.Handle("POST", "/gameplay", controller.GameplayRedirect)
	r.Handle("GET", "/gameplay/:console/:game", controller.Gameplay)
	r.Handle("GET", "/games", controller.ListGames)
	r.Handle("HEAD", "/download/game/:console/:rom", controller.GetRom)
	r.Handle("GET", "/download/game/:console/:rom", controller.GetRom)
	r.Handle("HEAD", "/download/bios/:console/:bios", controller.GetBios)
	r.Handle("GET", "/download/bios/:console/:bios", controller.GetBios)

	return r
}
