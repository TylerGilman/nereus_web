package handlers

import (
	"net/http"

	"nereus_web/views/about"
)

func HandleMin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, about.About(true))
}