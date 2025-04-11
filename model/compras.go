package model

import "time"

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

func (c *Compras) CalcularHorarioCompras() {
	c.Horario = time.Now()
}
