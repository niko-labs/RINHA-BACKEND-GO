package repositories

const (
	// TRANSACOES
	T_INSERT_INFO  = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, $3, $4);"
	TD_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'd', $3);"
	TC_STMT_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'c', $3);"

	// CREDITO - DEBITO
	CD_STMT_UPDATE = "UPDATE clientes SET saldo = $2 WHERE id = $1;"

	// EXTRATO - CONSULTA
	Q_EXTRATO = "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 10;"

	// CLIENTE - CONSULTA
	Q_CLIENTE_INFOS = "SELECT limite, saldo FROM clientes WHERE id = $1 LIMIT 1 FOR UPDATE;"

	// QUERY - GERA EXTRATO DE CLIENTE DIRETO PELO BANCO
	Q_EXTRATO_CLIENTE = `

	WITH ultimas_transacoes_cte AS (
		SELECT
				valor, tipo, descricao, realizada_em
		FROM transacoes
		WHERE
				cliente_id = $1
		ORDER BY realizada_em DESC
		LIMIT 10
	),
		cliente_info_cte AS (
		SELECT 
				saldo, limite
		FROM clientes
		WHERE
			id = $1
		LIMIT 1
	)

	SELECT
		jsonb_build_object(
			'saldo', jsonb_build_object(
				'total', (SELECT saldo FROM cliente_info_cte),
				'limite', (SELECT limite FROM cliente_info_cte),
				'data_extrato', now()::timestamp with time zone
			),
			'ultimas_transacoes',
				CASE
					WHEN COUNT(ultimas_transacoes_cte) = 0 THEN '[]'
				ELSE
				jsonb_agg(
					jsonb_build_object(
						'valor', ultimas_transacoes_cte.valor,
						'tipo', ultimas_transacoes_cte.tipo,
						'descricao', ultimas_transacoes_cte.descricao,
						'realizada_em', ultimas_transacoes_cte.realizada_em
					)
				)
			END
		)
		FROM
			ultimas_transacoes_cte;
  
	`
)
