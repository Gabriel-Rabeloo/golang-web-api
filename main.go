package main

import (
	"net/http"

	"github.com/Gabriel-Rabeloo/web-api/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
