package app

import (
	"encoding/base64"
	"fmt"
	"net/http"

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
	gc.Redirect(http.StatusFound, fmt.Sprintf("/gameplay/%s/%s", gc.PostFormMap("emulator"), gc.PostFormMap("rom")))
}

func (c *Controller) Gameplay(gc *gin.Context) {
	gameplay := c.service.GameplayDetail(gc.Param("console"), gc.Param("game"))
	gc.HTML(http.StatusOK, "gameplay.html", gin.H{"data": gameplay})
}

func (c *Controller) RomList(gc *gin.Context) {
	console := gc.DefaultQuery("console", "")
	roms := c.service.ListGames(console)
	gc.JSON(http.StatusOK, gin.H{"roms": roms})
}

func (c *Controller) RomDownload(gc *gin.Context) {
	path, _ := base64.StdEncoding.DecodeString(gc.Param("path"))
	url := string(path)

	resp, err := http.Get(url)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	if gc.Request.Method == "HEAD" {
		gc.Header("Content-Length", fmt.Sprintf("%d", resp.ContentLength))
		gc.Status(http.StatusOK)
		return
	}

	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func (c *Controller) BiosDownload(gc *gin.Context) {
	path, _ := base64.StdEncoding.DecodeString(gc.Param("path"))
	url := string(path)

	resp, err := http.Get(url)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	if gc.Request.Method == "HEAD" {
		gc.Header("Content-Length", fmt.Sprintf("%d", resp.ContentLength))
		gc.Status(http.StatusOK)
		return
	}

	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func NewController() *Controller {
	return &Controller{
		service: NewService(),
	}
}
