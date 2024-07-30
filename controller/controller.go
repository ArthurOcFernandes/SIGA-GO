package controller

import (
	"html/template"
	"net/http"
	"siga/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Students(w http.ResponseWriter, r *http.Request) {

	students := models.SearchStudentFromDataBase()

	temp.ExecuteTemplate(w, "StudentsList", students)
}

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Index", nil)
}
