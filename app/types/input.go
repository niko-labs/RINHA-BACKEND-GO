package types

type TransacaoInput struct {
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

const (
	KDEBITO                   = "d"
	KCREDITO                  = "c"
	KDESCRICAO_TAMANHO_MAXIMO = 10
	KDESCRICAO_VAZIA          = ""
)

func (t *TransacaoInput) Validar() bool {
	//  Tipo deve ser: "c" or "d"
	//  Descricao deve ser uma string com tamanho menor ou igual a 10 e não vazia ou ""
	//  Valor deve ser maior ou igual a 0

	return (t.Tipo == KCREDITO || t.Tipo == KDEBITO) && len(t.Descricao) <= KDESCRICAO_TAMANHO_MAXIMO && t.Descricao != KDESCRICAO_VAZIA
}
