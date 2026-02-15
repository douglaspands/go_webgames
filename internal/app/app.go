package app

import (
	"net/http"
	"os"
	controller "webgames/internal/controller"
	repository "webgames/internal/repository"
	service "webgames/internal/service"

	"github.com/gin-gonic/gin"
)

type App interface {
	Handler() http.Handler
}

type app struct {
	ginEngine          *gin.Engine
	gameplayController controller.GameplayController
}

func (a *app) Handler() http.Handler {
	if a.ginEngine == nil {
		a.setup()
	}
	return a.ginEngine.Handler()
}

func (a *app) setup() {
	mode := os.Getenv("GIN_MODE")
	if len(mode) == 0 {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	a.ginEngine = gin.Default()
	a.ginEngine.Use(gin.Recovery())

	a.ginEngine.LoadHTMLGlob("templates/*")
	a.ginEngine.Static("/static", "./static")
	a.ginEngine.StaticFile("/favicon.ico", "./static/favicon.ico")

	a.instance()
	a.router()
}

func (a *app) instance() {
	romRepository := repository.NewRomRepository()
	emulatorRepository := repository.NewEmulatorRepository()
	gamePlayService := service.NewGameplayService(emulatorRepository, romRepository)
	a.gameplayController = controller.NewGameplayController(gamePlayService)
}

func (a *app) router() {
	a.ginEngine.Handle("GET", "/", a.gameplayController.GetIndex)
	a.ginEngine.Handle("POST", "/gameplay", a.gameplayController.GameplayRedirect)
	a.ginEngine.Handle("GET", "/gameplay/:console/:game", a.gameplayController.Gameplay)
	a.ginEngine.Handle("GET", "/games", a.gameplayController.ListGames)
	a.ginEngine.Handle("HEAD", "/download/game/:console/:rom", a.gameplayController.GetRom)
	a.ginEngine.Handle("GET", "/download/game/:console/:rom", a.gameplayController.GetRom)
	a.ginEngine.Handle("HEAD", "/download/bios/:console/:bios", a.gameplayController.GetBios)
	a.ginEngine.Handle("GET", "/download/bios/:console/:bios", a.gameplayController.GetBios)
}

func NewApp() App {
	return &app{}
}
