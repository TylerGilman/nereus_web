package handlers

import (
	"net/http"

	"nereus_web/views/about"
)

type navigationData struct {
	page_number int
}

func HandleNavigation(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, about.About())
}
