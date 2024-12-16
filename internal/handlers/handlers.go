package handlers

import (
	"joshuamURD/testing/internal/renderer"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	renderer.Render(w, "index.page.html")
}