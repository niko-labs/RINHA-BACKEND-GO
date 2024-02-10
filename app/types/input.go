package types

type TransacaoInput struct {
	Valor     int64         `json:"valor"`
	Tipo      TipoTransacao `json:"tipo"`
	Descricao string        `json:"descricao"`
}
