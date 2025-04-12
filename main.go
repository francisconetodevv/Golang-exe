package main

import (
	"exercicioOne/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Testando...")

	/*Inicializando o slice*/
	var nomeItens []string
	nomeItens = append(nomeItens, "Arroz")
	nomeItens = append(nomeItens, "Feijao")
	nomeItens = append(nomeItens, "Carne")
	nomeItens = append(nomeItens, "Macarrão")

	compras, err := model.NewCompra("Mercadinho Jesus", time.Now(), nomeItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compras)
	}

}
