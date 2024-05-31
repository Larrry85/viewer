// NEW VERSION

// apiServer.go
package gofiles

import (
	"net/http"
	"path/filepath"
	"mime"
)


// Handles image serving with CORS support.
// CorsMiddleware adds CORS headers to responses.
// Servereiden pitäisi tehdä yhteistyötä toistensa kanssa
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from localhost:8080
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request, return early with status code 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// ServeImage serves image files from the api/img directory
// serve images
func ServeImage(w http.ResponseWriter, r *http.Request) {
	// Join the directory with the requested file path
	filePath := filepath.Join("api", "img", filepath.Base(r.URL.Path))
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
	http.ServeFile(w, r, filePath)
}