package handlers

import (
	"net/http"

	"nereus_web/views/blog"
)

func HandleBlogIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, blog.Blog())
}
