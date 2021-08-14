package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1qaz!QAZ"
	dbname   = "neoway"
  )

//Conectar abre a conexão com o banco de dados a a retorna.
func Conectar() (*sql.DB, error) {
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)


	db, erro := sql.Open("postgres", psqlConnectionString)

	if erro != nil {
		return nil , erro
	}

	if db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}