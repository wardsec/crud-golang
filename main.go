package main

import (
	"net/http"
	"wardsec/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)

}
