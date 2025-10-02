package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tiilyat/embed-go-front/ui"
)

func NewServer(config ServeConfig) *gin.Engine {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	bindRoutes(router, BindRoutesConfig{
		DevMode:           config.UIDevMode,
		FrontDevServerURL: config.UIDevServerURL,
		DistDirFS:         ui.DistDirFS,
	})

	return router
}
