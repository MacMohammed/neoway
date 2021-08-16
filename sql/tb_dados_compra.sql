create table if not exists tb_dados_compra (
	cpf double precision not null,
	private int,
	incompleto int,
	data_ultima_compra date null,
	ticket_medio numeric null,
	ticket_ultima_compra numeric null,
	loja_mais_frequente double precision null,
	loja_ultima_compra double precision null
);