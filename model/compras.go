/*Estudo sobre Golang*/

package model

import (
	"errors"
	"time"
)

type Compras struct {
	NomeMercado string
	Horario     time.Time
	Itens       []ItemCompras
	/* Propriedade que será um slice */
	/* Isso aqui define que os itens serão uma lista*/
	/* Faz jus, pos é uma lista de itens */
}

type ItemCompras struct {
	NomeItem string
}

/* Método para realizar compra */
/* Passamos os parâmetros da compra que são os itens, o retorno será a compra e um erro, caso ocorra.*/
/* É preciso retornar o erro também, pois podemos tratar caso o usuário não passar nenhum valor */
func NewCompra(mercado string, data time.Time, nomeItens []string) (*Compras, error) {

	/*Validando se as informações repassadas para função se são nulas ou não*/
	if mercado == "" {
		return nil, errors.New("Mercado é obrigatório")
	}

	if len(nomeItens) == 0 {
		return nil, errors.New("Itens são obrigatórios")
	}

	var itens []ItemCompras

	for _, nome := range nomeItens {
		itens = append(itens, ItemCompras{NomeItem: nome})
	}

	return &Compras{
		NomeMercado: mercado,
		Horario:     time.Now(),
		Itens:       itens,
	}, nil
}
