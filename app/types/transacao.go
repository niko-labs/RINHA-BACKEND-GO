package types

import "time"

type TbTransacao struct {
	Id          int16     `json:"id" omitempty:"true"`
	ClienteId   int16     `json:"cliente_id" omitempty:"true"`
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadoEm time.Time `json:"realizado_em"`
}

type TbTransacaoApenasId struct {
	Id int8 `json:"id"`
}
