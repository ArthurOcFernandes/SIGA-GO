package main

import (
	"net/http"
	"siga/routes"
)

func main() {

	routes.LoadRoutes()

	_ = http.ListenAndServe(":8000", nil)
}
