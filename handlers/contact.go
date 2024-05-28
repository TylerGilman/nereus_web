package handlers

import (
	"net/http"

	"nereus_web/views/contact"
)

func HandleContactIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, contact.Contact(false))
}
