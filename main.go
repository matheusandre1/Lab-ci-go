package main

import (
	"github.com/matheusandre1/Lab-ci-go/database"
	"github.com/matheusandre1/Lab-ci-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
