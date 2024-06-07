// serveImg.go
package gofiles

import (
	"net/http"			// handling HTTP requests and responses
	"path/filepath"		// manipulating file paths
	"mime"				// determining MIME types based on file extensions
)


// adds CORS headers to HTTP responses to allow allow cross-origin requests
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allows requests from http://localhost:8080
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		// specifies the allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// specifies the allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// if a preflight request, it responds with 200 OK and returns early
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK) // 200
			return
		}
		// for other request methods, it calls the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// serves image files from the api/img directory
func ServeImage(w http.ResponseWriter, r *http.Request) {
	// constructs file path by joining the api/img directory with the base name of the requested URL path
	filePath := filepath.Join("api", "img", filepath.Base(r.URL.Path))
	ext := filepath.Ext(filePath) // gets the file extension
	mimeType := mime.TypeByExtension(ext) // determines the MIME type based on the file extension
	// if a MIME type is found, it sets the Content-Type header.
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	} // serves the file at the constructed file path
	http.ServeFile(w, r, filePath)
}