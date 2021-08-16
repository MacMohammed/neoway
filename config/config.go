package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


var (
	//PsqlConnectionString é a string de conexão com o servidor Postgres
	PsqlConnectionString = ""

	//Porta é a porta definida onde a aplicação estará rodando
	Porta = 0

)

//Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	PsqlConnectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", 
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

}