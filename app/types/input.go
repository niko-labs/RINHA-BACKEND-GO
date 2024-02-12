package types

type TransacaoInput struct {
	Valor     int64         `json:"valor"`
	Tipo      TipoTransacao `json:"tipo"`
	Descricao string        `json:"descricao"`
}

func (t *TransacaoInput) Validar() bool {
	//  Tipo deve ser: "c" or "d"
	//  Descricao deve ser uma string com tamanho menor ou igual a 10 e não vazia ou ""
	//  Valor deve ser maior ou igual a 0

	return (t.Tipo == "c" || t.Tipo == "d") && len(t.Descricao) <= 10 && t.Descricao != "" && t.Valor >= 0
}
