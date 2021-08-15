package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "bd_postgres"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "neoway"
  )

//Conectar abre a conexão com o banco de dados a a retorna.
func Conectar() (*sql.DB, error) {
	psqlConnectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	db, erro := sql.Open("postgres", psqlConnectionString)

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