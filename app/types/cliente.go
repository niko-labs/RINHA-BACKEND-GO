package types

type TbCliente struct {
	Id     int16  `json:"id"`
	Nome   string `json:"nome"`
	Limite int64  `json:"limite"`
	Saldo  int64  `json:"saldo"`
}

type TbClienteApenasId struct {
	Id int8 `json:"id"`
}
