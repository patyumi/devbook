package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
