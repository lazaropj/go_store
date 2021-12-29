package main

import (
	"net/http"

	"github.com/lazaropj/go_store/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
