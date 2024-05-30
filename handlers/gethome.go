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
	isMinimized := r.URL.Query().Get("minimize")
	// If header = content render content
	// else render whole screen
	if tab == "" {
		tab = "About"
	}
	if isMinimized == "" {
		isMinimized = "false"
	}
	r = r.WithContext(context.WithValue(r.Context(), "tab", tab))
	r = r.WithContext(context.WithValue(r.Context(), "minimize", isMinimized))

	fmt.Println("ctx.tab: ", tab)
	fmt.Println("ctx.isMinimized ", isMinimized)
	header := r.Header.Get("View")
	if header == "content" {
		minimize := true
		if isMinimized == "false" {
			minimize = false
		}
		if tab == "About" {
			Render(w, r, components.Navigation(minimize, 0))
			return Render(w, r, about.About())
		}
		if tab == "Contact" {
			Render(w, r, components.Navigation(minimize, 1))
			return Render(w, r, contact.Contact())
		}
		if tab == "Blog" {
			Render(w, r, components.Navigation(minimize, 2))

			return Render(w, r, blog.Blog())
		}
	}
	return Render(w, r, home.Index())
}
