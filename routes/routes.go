package routes

import (
	"net/http"
	"siga/controller"
)

func LoadRoutes() {
	http.HandleFunc("/Students", controller.Students)
}
