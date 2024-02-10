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
	VALUES (
	        'o barato sai caro', 1000 * 100, 0
	    ),
	    ('zan corp ltda', 800 * 100, 0),
	    ('les cruders', 10000 * 100, 0),
	    (
	        'padaria joia de cocaia', 100000 * 100, 0
	    ),
	    ('kid mais', 5000 * 100, 0);
END;
$$; 

CREATE INDEX idx_compound_cliente_id_realizado_em ON transacoes (cliente_id, realizada_em);

CREATE INDEX idx_cliente_id ON transacoes (cliente_id);