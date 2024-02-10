package database

const (
	// TRANSACOES
	TD_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'd', $3);"
	TC_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'c', $3);"

	// CREDITO - DEBITO
	CD_STMT_UPDATE = "UPDATE clientes SET saldo = $1 WHERE id = $2;"
)
