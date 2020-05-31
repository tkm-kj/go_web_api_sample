package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tkm-kj/go_web_api_sample/internal/routes"
)

// GAE用のサーバー
func main() {
	e := echo.New()

	routes.Register(e)
	http.Handle("/", e)

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
