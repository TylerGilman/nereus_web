package handlers

import (
	"fmt"
	"net/http"

	"nereus_web/views/about"
)

func HandlePage(w http.ResponseWriter, r *http.Request) error {
	page := r.URL.Query().Get("page")
	mode := r.URL.Query().Get("mode")
	minimized := r.URL.Query().Get("minimized")
	fmt.Println(page)
	fmt.Println(mode)
	fmt.Println(minimized)

	return Render(w, r, about.About())
}
