package spafileserver

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/jordan-wright/unindexed"
)

func Handler(rootPath, directoryPath string) http.Handler {
	return http.StripPrefix(rootPath,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, err := os.Stat(filepath.Join(directoryPath, filepath.Clean(r.URL.Path)))
				if err != nil {
					http.ServeFile(w, r, filepath.Join(directoryPath, "index.html"))
					return
				}
				http.FileServer(unindexed.Dir(directoryPath)).ServeHTTP(w, r)
			},
		),
	)
}
