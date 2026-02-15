package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	service "webgames/internal/service"

	"github.com/gin-gonic/gin"
)

type GameplayController interface {
	GetIndex(gc *gin.Context)
	GameplayRedirect(gc *gin.Context)
	Gameplay(gc *gin.Context)
	ListGames(gc *gin.Context)
	GetRom(gc *gin.Context)
	GetBios(gc *gin.Context)
}

type gameplayController struct {
	gameplayService service.GameplayService
}

func (c gameplayController) allowSharedArrayBuffer(gc *gin.Context) {
	gc.Header("Cross-Origin-Opener-Policy", "same-origin")
	gc.Header("Cross-Origin-Embedder-Policy", "require-corp")
}

func (c gameplayController) responseHeadHttpMethod(gc *gin.Context, ContentLength int64) {
	gc.Header("Content-Type", "application/octet-stream")
	gc.Header("Content-Length", strconv.FormatInt(ContentLength, 10))
	gc.Status(http.StatusOK)
}

func (c *gameplayController) GetIndex(gc *gin.Context) {
	emulators := c.gameplayService.ListConsoles()
	gc.HTML(http.StatusOK, "index.html", gin.H{"data": emulators})
}

func (c *gameplayController) GameplayRedirect(gc *gin.Context) {
	console := gc.PostForm("console")
	game := gc.PostForm("game")
	gc.Redirect(http.StatusFound, fmt.Sprintf("/gameplay/%s/%s", url.PathEscape(console), url.PathEscape(game)))
}

func (c *gameplayController) Gameplay(gc *gin.Context) {
	console := gc.Param("console")
	game := gc.Param("game")
	gameplay := c.gameplayService.GameplayDetail(console, game)
	if gameplay.Threads == true {
		c.allowSharedArrayBuffer(gc)
	}
	gc.HTML(http.StatusOK, "gameplay.html", gin.H{"data": gameplay})
}

func (c *gameplayController) ListGames(gc *gin.Context) {
	console := gc.DefaultQuery("console", "")
	games := c.gameplayService.ListGames(console)
	gc.JSON(http.StatusOK, gin.H{"data": games})
}

func (c *gameplayController) GetRom(gc *gin.Context) {
	console := gc.Param("console")
	rom := gc.Param("rom")
	game := strings.TrimSuffix(rom, filepath.Ext(rom))
	gameplay := c.gameplayService.GameplayDetail(console, game)
	if gc.Request.Method == "HEAD" {
		resp, err := http.Head(gameplay.RomUrl)
		if err != nil {
			gc.Error(err)
			return
		}
		c.responseHeadHttpMethod(gc, resp.ContentLength)
	} else {
		resp, err := http.Get(gameplay.RomUrl)
		if err != nil {
			gc.Error(err)
			return
		}
		defer resp.Body.Close()
		gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
	}
}

func (c *gameplayController) GetBios(gc *gin.Context) {
	console := gc.Param("console")
	emulator := c.gameplayService.GetConsole(console)
	if gc.Request.Method == "HEAD" {
		resp, err := http.Head(emulator.BiosUrl)
		if err != nil {
			gc.Error(err)
			return
		}
		c.responseHeadHttpMethod(gc, resp.ContentLength)
	} else {
		resp, err := http.Get(emulator.BiosUrl)
		if err != nil {
			gc.Error(err)
			return
		}
		defer resp.Body.Close()
		gc.DataFromReader(http.StatusOK, resp.ContentLength, "application/octet-stream", resp.Body, map[string]string{})
	}
}

func NewGameplayController(gameplayService service.GameplayService) GameplayController {
	return &gameplayController{
		gameplayService: gameplayService,
	}
}
