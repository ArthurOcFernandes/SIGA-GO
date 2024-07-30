package routes

import (
	"net/http"
	"siga/controller"
)

func LoadRoutes() {
	http.HandleFunc("/Students", controller.Students)
	http.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("./templates/css"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("./templates/img"))))
	http.HandleFunc("/", controller.Index)

}
