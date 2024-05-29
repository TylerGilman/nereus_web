package handlers

import (
	"net/http"

	"nereus_web/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index("about", "dark", false))
}
