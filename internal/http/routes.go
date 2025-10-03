package http

import (
	"io/fs"

	"github.com/gin-gonic/gin"
)

type BindRoutesConfig struct {
	DevMode           bool
	FrontDevServerURL string
	DistDirFS         fs.FS
}

func bindRoutes(router *gin.Engine, config BindRoutesConfig) {
	api := router.Group("/api")
	{
		api.GET("/health", healthHandler())
	}

	router.NoRoute(staticHandler(config.DevMode, config.FrontDevServerURL, config.DistDirFS))
}
