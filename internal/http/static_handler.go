package http

import (
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func serveStaticFiles(fsys fs.FS, indexFallback bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := filepath.ToSlash(filepath.Clean(c.Request.URL.Path))
		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}

		_, err := fs.Stat(fsys, path)
		if err != nil {
			// File doesn't exist
			if indexFallback && path != "index.html" {
				// Serve index.html for SPA routing
				http.ServeFileFS(c.Writer, c.Request, fsys, "index.html")
				return
			}
			c.Status(http.StatusNotFound)
			return
		}

		// File exists, serve it
		http.ServeFileFS(c.Writer, c.Request, fsys, path)
	}
}

func serveDevProxy(frontDevServerURL string) gin.HandlerFunc {
	target, err := url.Parse(frontDevServerURL)
	if err != nil {
		log.Fatalf("Failed to parse Vite URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func staticHandler(devMode bool, frontDevServerURL string, distDirFS fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		if devMode {
			serveDevProxy(frontDevServerURL)(c)
		} else {
			serveStaticFiles(distDirFS, true)(c)
		}
	}
}
