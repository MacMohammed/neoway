create table if not exists tb_dados_compra (
	cpf double precision not null,
	private int,
	incompleto int,
	data_ultima_compra date,
	ticket_medio numeric,
	ticket_ultima_compra numeric,
	loja_mais_frequente double precision null,
	loja_ultima_compra double precision null
);