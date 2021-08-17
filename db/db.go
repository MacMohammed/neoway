package db

import (
	"database/sql"
	"fmt"

	"neoway/config"

	_ "github.com/lib/pq"
)

//Conectar abre a conexão com o banco de dados retorna. Se algo der errado, um erro será retornado.
func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("postgres", config.PsqlConnectionString)

	if erro != nil {
		return nil , erro
	}

	if db.Ping(); erro != nil {
		db.Close()
		fmt.Println("Não conenctou no banco...")
		return nil, erro
	}

	return db, nil
}