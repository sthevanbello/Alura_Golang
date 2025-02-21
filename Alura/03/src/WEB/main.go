package main

import (
	"WEB/produto"
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectarBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// produtos := []produto.Produto{
	// 	{Name: "Camiseta", Description: "Camiseta preta", Quantity: 20, Price: 59.99},
	// 	{Name: "Camiseta", Description: "Camiseta Azul", Quantity: 25, Price: 49.99},
	// }
	db := conectarBancoDeDados()
	selectTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := produto.Produto{}
	produtos := []produto.Produto{}
	for selectTodosProdutos.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectTodosProdutos.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		produtos = append(produtos, p)
	}
	templates.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}

func conectarBancoDeDados() *sql.DB {
	connection := "user=admin password=senha123 dbname=DbGo host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
