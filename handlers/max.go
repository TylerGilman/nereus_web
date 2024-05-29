package handlers

import (
	"net/http"
)

func HandleMax(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, nil)
}
