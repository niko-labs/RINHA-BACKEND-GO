CREATE UNLOGGED
TABLE clientes (
    id SERIAL PRIMARY KEY, nome VARCHAR(50) NOT NULL, limite INTEGER NOT NULL, saldo INTEGER DEFAULT 0
);

CREATE UNLOGGED
TABLE transacoes (
    id SERIAL PRIMARY KEY, cliente_id INTEGER NOT NULL, valor INTEGER NOT NULL, tipo CHAR(1) NOT NULL, descricao VARCHAR(10) NOT NULL, realizada_em TIMESTAMP NOT NULL DEFAULT NOW(), CONSTRAINT fk_clientes_transacoes_id FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

DO $$ 
BEGIN 
	INSERT INTO
	    clientes (nome, limite, saldo)
	VALUES ('Asuka', 1000 * 100, 0),
	    ('Rin', 800 * 100, 0),
	    ('Shinji', 10000 * 100, 0),
	    ('Fern', 100000 * 100, 0),
	    ('Frienren', 5000 * 100, 0);
END;
$$; 

-- Map (
--     "id" -> 1, "limite" -> 1000 * 100
-- ),
-- Map (
--     "id" -> 2, "limite" -> 800 * 100
-- ),
-- Map (
--     "id" -> 3, "limite" -> 10000 * 100
-- ),
-- Map (
--     "id" -> 4, "limite" -> 100000 * 100
-- ),
-- Map (
--     "id" -> 5, "limite" -> 5000 * 100
-- ),

CREATE INDEX idx_compound_cliente_id_realizado_em ON transacoes (cliente_id, realizada_em);

CREATE INDEX idx_transacoes_cliente_id ON transacoes (cliente_id);

CREATE INDEX idx_clientes_id ON clientes (id);