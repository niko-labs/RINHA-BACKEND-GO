CREATE UNLOGGED
TABLE clientes (
    id SERIAL PRIMARY KEY, nome VARCHAR(50) NOT NULL, limite INTEGER NOT NULL
);

CREATE UNLOGGED
TABLE transacoes (
    id SERIAL PRIMARY KEY, cliente_id INTEGER NOT NULL, valor INTEGER NOT NULL, tipo CHAR(1) NOT NULL, descricao VARCHAR(10) NOT NULL, realizada_em TIMESTAMP NOT NULL DEFAULT NOW(), CONSTRAINT fk_clientes_transacoes_id FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

CREATE UNLOGGED
TABLE saldos (
    id SERIAL PRIMARY KEY, cliente_id INTEGER NOT NULL, valor INTEGER NOT NULL DEFAULT 0, CONSTRAINT fk_clientes_saldos_id FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

DO $$ 
BEGIN 
	INSERT INTO
	    clientes (nome, limite)
	VALUES ('Asuka', 1000 * 100),
	    ('Rin', 5000 * 100),
	    ('Shinji', 800 * 100),
	    ('Fern', 10000 * 100),
	    ('Frienren', 100000 * 100);
	INSERT INTO
	    saldos (cliente_id, valor)
	SELECT id, 0
	FROM clientes;
END;
$$; 

CREATE INDEX idx_compound_cliente_id_realizado_em ON transacoes (cliente_id, realizada_em);

CREATE INDEX idx_transacoes_cliente_id ON transacoes (cliente_id);

CREATE INDEX idx_clientes_id ON clientes (id);

-- --- CONVERTER TO ME TABLES
-- SELECT balances.amount AS amount, accounts.limit_amount AS limit_amount
-- FROM accounts
--     JOIN balances ON balances.account_id = accounts.id
-- WHERE
--     accounts.id = $1 FOR
-- UPDATE
-- --- CONVERTER TO ME TABLES
-- SELECT saldos.valor, clientes.limite
-- FROM saldos
--     JOIN clientes ON clientes.id = saldos.cliente_id
-- WHERE
--     saldos.cliente_id = $1 FOR
-- UPDATE