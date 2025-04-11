package main

import (
	"exercicioOne/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Testando...")

	var items []model.ItemCompras
	items = append(items, model.ItemCompras{NomeItem: "Arroz"})
	items = append(items, model.ItemCompras{NomeItem: "Feijao"})

	compras := model.Compras{
		NomeMercado: "Americanas",
		Horario:     time.Now(),
		Itens:       items,
	}

	fmt.Println(compras)
}
