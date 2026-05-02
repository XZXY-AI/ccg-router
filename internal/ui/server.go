// Package ui serves the read-only local web panel.
package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:assets
var assets embed.FS

func Handler() http.Handler {
	sub, _ := fs.Sub(assets, "assets")
	return http.StripPrefix("/ui/", http.FileServer(http.FS(sub)))
}
