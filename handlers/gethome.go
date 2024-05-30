package handlers

import (
	"context"
	"fmt"
	"net/http"

	"nereus_web/views/about"
	"nereus_web/views/blog"
	"nereus_web/views/components"
	"nereus_web/views/contact"
	"nereus_web/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	tab := r.URL.Query().Get("tab")
	// If header = content render content
	// else render whole screen
	if tab == "" {
		tab = "About"
	}
	r = r.WithContext(context.WithValue(r.Context(), "tab", tab))

	fmt.Println("Context ", tab)
	header := r.Header.Get("View")
	if header == "content" {
		if tab == "About" {
			Render(w, r, components.Navigation(false, 0))
			return Render(w, r, about.About())
		}
		if tab == "Contact" {
			Render(w, r, components.Navigation(false, 1))
			return Render(w, r, contact.Contact())
		}
		if tab == "Blog" {
			Render(w, r, components.Navigation(false, 2))
			return Render(w, r, blog.Blog())
		}
	}
	return Render(w, r, home.Index())
}