package http

import (
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
)

func serveStaticFiles(fsys fs.FS, indexFallback bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.ToSlash(filepath.Clean(r.URL.Path))
		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}

		_, err := fs.Stat(fsys, path)
		if err != nil {
			// File doesn't exist
			if indexFallback && path != "index.html" {
				// Serve index.html for SPA routing
				http.ServeFileFS(w, r, fsys, "index.html")
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// File exists, serve it
		http.ServeFileFS(w, r, fsys, path)
	}
}

func serveDevProxy(frontDevServerURL string) http.HandlerFunc {
	target, err := url.Parse(frontDevServerURL)
	if err != nil {
		log.Fatalf("Failed to parse Vite URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func staticHandler(devMode bool, frontDevServerURL string, distDirFS fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if devMode {
			serveDevProxy(frontDevServerURL)(w, r)
		} else {
			serveStaticFiles(distDirFS, true)(w, r)
		}
	}
}
