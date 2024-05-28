package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"nereus_web/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/min", handlers.Make(handlers.HandleMin))
	router.Get("/max", handlers.Make(handlers.HandleMax))
	router.Get("/about", handlers.Make(handlers.HandleAboutIndex))
	router.Get("/blog", handlers.Make(handlers.HandleBlogIndex))
	router.Get("/contact", handlers.Make(handlers.HandleContactIndex))
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
