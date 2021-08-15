package model

import (
	"bufio"
	"errors"
	"fmt"
	db "neoway/db"
	"neoway/utils"
	"os"
	"strings"
	"time"
)

//dadosCompra representa as colunas do arquivo a ser tratado
type dadosCompra struct {
	Cpf                string
	Private            string
	Incompleto         string
	DataUltimaCompra   *string
	TicketMedio        string
	TicketUltimaCompra string
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}

//PersistData persiste no banco os dados do arquivo
func PersistData(file string)  error {
	fmt.Printf("Início do processamento: %v\n", time.Now().Format("2006-01-02 15:04:05"))

	arquivo, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	//Esse trecho do código garante que a primeira linha do arquivo será ignorada
	scanner.Scan()

	db, err := db.Conectar();
	if err != nil {

		fmt.Println("O erro foi no acesso ao banco")
		return err
	}

	defer db.Close()

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if len(s) != 8 {
			fmt.Printf("Fim do processamento: %v\n", time.Now().Format("2006-01-02 15:04:05"))
			return errors.New("arquivo não compatível.")
		}



		cpf, err := utils.LimparCpfCnpj(s[0])
		if err != nil {
			fmt.Println(err)
		}

		lojaMaisFrequente, err := utils.LimparCpfCnpj(s[6])
		if err != nil {
			fmt.Println(err)
		}

		lojaUltimaCompra, err := utils.LimparCpfCnpj(s[7])
		if err != nil {
			fmt.Println(err)
		}

		TicketMedio, err := utils.RemoveVirgula(strings.TrimSpace(s[4]))
		if err != nil {
			fmt.Println(err)
		}

		TicketUltimaCompra, err := utils.RemoveVirgula(strings.TrimSpace(s[5]))
		if err != nil {
			fmt.Println(err)
		}

		DataUltimaCompra, err := utils.RemoveNull(strings.TrimSpace(s[3]))
		if err != nil {
			fmt.Println(err)
		}

		d := dadosCompra {
			Cpf: cpf,
			Private: strings.TrimSpace(s[1]),
			Incompleto: strings.TrimSpace(s[2]),
			DataUltimaCompra: DataUltimaCompra,
			TicketMedio: TicketMedio,
			TicketUltimaCompra: TicketUltimaCompra,
			LojaMaisFrequente: lojaMaisFrequente,
			LojaUltimaCompra: lojaUltimaCompra,
		}

		statement, err := db.Prepare("insert into tb_dados_compra (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) values ($1, $2, $3, $4, $5, $6, $7, $8);",)
		if err != nil {
			fmt.Println(err)
			return err
		}

		defer statement.Close();

		if _, err := statement.Exec(d.Cpf, d.Private, d.Incompleto, d.DataUltimaCompra, d.TicketMedio, d.TicketUltimaCompra, d.LojaMaisFrequente, d.LojaUltimaCompra,); err != nil {
			fmt.Println(err)
			return err
		}
	}

	fmt.Printf("Fim do processamento: %v\n", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}