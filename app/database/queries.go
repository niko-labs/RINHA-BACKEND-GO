package database

const (
	T_DC_CTE_UNICO = `
	WITH
		cte_atualizar_saldo AS (
			UPDATE clientes AS clt
				SET saldo = CASE  
						WHEN $3='c' THEN (saldo + $2)
						WHEN $3='d' AND (saldo - $2) > -limite THEN (saldo - $2) 
						
						ELSE saldo
					END
			WHERE clt.id = $1
			RETURNING saldo, limite
		),
		cte_inserir_transacao AS (
			INSERT INTO transacoes (cliente_id, valor, tipo, descricao)
			VALUES ($1, $2, $3, $4)
		)
	SELECT * FROM cte_atualizar_saldo;
	`

	// TRANSACOES
	T_DC_INSERT = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, 'd', $3);"
	T_DC_UPDATE = "UPDATE clientes SET saldo = $1 WHERE id = $2;"

	// CLIENTE - CONSULTA
	Q_CLIENTE = "SELECT id, limite, saldo FROM clientes WHERE id = $1 LIMIT 1"

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
