package http

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/tiilyat/embed-go-front/ui"
)

// serveStaticFiles serves files from the embedded filesystem.
// If indexFallback is true, serves index.html for non-existent paths (SPA routing).
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

// spaHandler serves embedded static files with SPA fallback routing.
func spaHandler() http.HandlerFunc {
	return serveStaticFiles(ui.DistDirFS, true)
}
