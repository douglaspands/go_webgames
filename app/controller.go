package app

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func (c *Controller) GetIndex(gc *gin.Context) {
	emulators := c.service.ListConsoles()
	gc.HTML(http.StatusOK, "index.html", gin.H{"data": emulators})
}

func (c *Controller) GameplayRedirect(gc *gin.Context) {
	emulator := gc.PostForm("emulator")
	rom := gc.PostForm("rom")
	gc.Redirect(http.StatusFound, fmt.Sprintf("/gameplay/%s/%s", url.PathEscape(emulator), url.PathEscape(rom)))
}

func (c *Controller) Gameplay(gc *gin.Context) {
	console := gc.Param("console")
	game := gc.Param("game")
	gameplay := c.service.GameplayDetail(console, game)
	gc.HTML(http.StatusOK, "gameplay.html", gin.H{"data": gameplay})
}

func (c *Controller) RomList(gc *gin.Context) {
	console := gc.DefaultQuery("console", "")
	roms := c.service.ListGames(console)
	gc.JSON(http.StatusOK, gin.H{"data": roms})
}

func (c *Controller) RomDownload(gc *gin.Context) {
	if gc.Request.Method == "HEAD" {
		gc.Status(http.StatusOK)
		return
	}
	path := gc.Param("path")
	bpath, _ := base64.StdEncoding.DecodeString(path)
	url := string(bpath)

	resp, err := http.Get(url)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func (c *Controller) BiosDownload(gc *gin.Context) {
	if gc.Request.Method == "HEAD" {
		gc.Status(http.StatusOK)
		return
	}
	path, _ := base64.StdEncoding.DecodeString(gc.Param("path"))
	url := string(path)

	resp, err := http.Get(url)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func NewController() *Controller {
	return &Controller{
		service: NewService(),
	}
}
