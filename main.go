package main

import (
	"net/http"

	"loja.go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
