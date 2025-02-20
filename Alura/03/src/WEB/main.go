package main

import (
	"WEB/produto"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []produto.Produto{
		{Name: "Camiseta", Description: "Camiseta preta", Quantity: 20, Price: 59.99},
		{Name: "Camiseta", Description: "Camiseta Azul", Quantity: 25, Price: 49.99},
	}

	templates.ExecuteTemplate(w, "index", produtos)
}
