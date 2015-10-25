package static

import (
	"net/http"
	"os"
)

// Handler return http hander serving static content
func Handler(rootPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath := r.URL.Path[len("/"):]

		file, err := os.Open(rootPath + requestedPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		stat, err := file.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, requestedPath, stat.ModTime(), file)
	})
}
