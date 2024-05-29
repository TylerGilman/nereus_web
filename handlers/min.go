package handlers

import (
	"net/http"
)

func HandleMin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, nil)
}
