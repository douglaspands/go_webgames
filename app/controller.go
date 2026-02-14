package app

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func (c Controller) allowSharedArrayBuffer(gc *gin.Context) {
	gc.Header("Cross-Origin-Opener-Policy", "same-origin")
	gc.Header("Cross-Origin-Embedder-Policy", "require-corp")
}

func (c Controller) responseHeadHttpMethod(gc *gin.Context, ContentLength int64) {
	gc.Header("Content-Type", "application/octet-stream")
	gc.Header("Content-Length", strconv.FormatInt(ContentLength, 10))
	gc.Status(http.StatusOK)
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
	if gameplay.Threads == true {
		c.allowSharedArrayBuffer(gc)
	}
	gc.HTML(http.StatusOK, "gameplay.html", gin.H{"data": gameplay})
}

func (c *Controller) ListGames(gc *gin.Context) {
	console := gc.DefaultQuery("console", "")
	games := c.service.ListGames(console)
	gc.JSON(http.StatusOK, gin.H{"data": games})
}

func (c *Controller) GetRom(gc *gin.Context) {
	console := gc.Param("console")
	rom := gc.Param("rom")
	game := strings.TrimSuffix(rom, filepath.Ext(rom))

	gameplay := c.service.GameplayDetail(console, game)

	resp, err := http.Get(gameplay.RomUrl)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	if gc.Request.Method == "HEAD" {
		c.responseHeadHttpMethod(gc, resp.ContentLength)
		return
	}
	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func (c *Controller) GetBios(gc *gin.Context) {
	console := gc.Param("console")
	emulator := c.service.GetConsole(console)

	resp, err := http.Get(emulator.BiosUrl)
	if err != nil {
		gc.Error(err)
		return
	}
	defer resp.Body.Close()

	if gc.Request.Method == "HEAD" {
		c.responseHeadHttpMethod(gc, resp.ContentLength)
		return
	}
	gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
}

func NewController() *Controller {
	return &Controller{
		service: NewService(),
	}
}
