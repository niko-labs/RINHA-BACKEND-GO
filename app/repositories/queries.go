package repositories

const (
	// TRANSACOES
	T_INSERT_INFO = "INSERT INTO transacoes (cliente_id, valor, tipo, descricao) VALUES ($1, $2, $3, $4);"

	UPDATE_SALDO_D = "UPDATE saldos SET valor = valor - $2 WHERE cliente_id = $1"
	UPDATE_SALDO_C = "UPDATE saldos SET valor = valor + $2 WHERE cliente_id = $1"

	// CLIENTE - CONSULTA
	CLIENTE_INFO = "SELECT saldos.valor, clientes.limite FROM saldos JOIN clientes ON clientes.id = saldos.cliente_id WHERE saldos.cliente_id = $1 FOR UPDATE"

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
		--- FOR UPDATE
	),
		cliente_info_cte AS (
		SELECT 
				valor, limite
		FROM clientes
		JOIN saldos ON saldos.cliente_id = clientes.id
		WHERE
				clientes.id = $1
		LIMIT 1 FOR UPDATE
	)

	SELECT
		jsonb_build_object(
			'saldo', jsonb_build_object(
				'total', (SELECT valor FROM cliente_info_cte),
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

	CTE_CDI = `
	WITH
		cte_atualizar_saldo AS (
			UPDATE saldos AS sd
				SET valor = CASE  
						WHEN $3='c' THEN (valor + $2)
						WHEN $3='d' AND (valor - $2) < -cl.limite THEN (valor - $2) 
						ELSE valor
					END
				FROM clientes cl
				WHERE sd.cliente_id = cl.id AND cl.id = $1
				RETURNING sd.valor, cl.limite
		),
		cte_inserir_transacao AS (
			INSERT INTO transacoes (cliente_id, valor, tipo, descricao)
			VALUES ($1, $2, $3, $4)
		)
	SELECT * FROM cte_atualizar_saldo FOR UPDATE;
	`
)
