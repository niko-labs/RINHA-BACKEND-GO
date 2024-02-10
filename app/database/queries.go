package database

const (
	// TRANSACOES
	TD_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'd', $3);"
	TC_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'c', $3);"

	// CREDITO - DEBITO
	CD_STMT_UPDATE = "UPDATE clientes SET saldo = $1 WHERE id = $2;"

	// EXTRATO - CONSULTA
	Q_EXTRATO = "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 10;"

	// CLIENTE - CONSULTA
	Q_CLIENTE = "SELECT id, limite, saldo FROM clientes WHERE id = $1 LIMIT 1"
)
