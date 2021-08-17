package model

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"neoway/db"
	"neoway/utils"
	"os"
	"strconv"
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
func PersistData(file string) error  {
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


	// db, err := db.Conectar();
	// if err != nil {

	// 	fmt.Println("O erro foi no acesso ao banco")
	// 	return err
	// }
	// defer db.Close()

	dados := []*dadosCompra{}

	//Aqui é feito um loop para tratar e persistir as linhas do arquivo no banco de dados.
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

		dados = append(dados, &d)
	}

	//Variável temporária para persistir os dados no banco
	temp := []*dadosCompra{}

	for _, dd := range dados {
		temp = append(temp, dd)

		//Afim de garantir performance, os dados são inseridos no banco a cada 5000 linhas
		if len(temp) == 5000 {
			err = persistir(temp)
			if err != nil {
				return err
			} else {
				temp = []*dadosCompra{}
			}
		}
	}

	//Caso fique um saldo, essa parte do código garante a persistência do saldo.
	if len(temp) > 0 {
		err = persistir(temp)
		if err != nil {
			return err
		} else {
			temp = []*dadosCompra{}
		}
	}

	fmt.Printf("Fim do processamento: %v\n", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}

func persistir(dados []*dadosCompra) error {

	query := `insert into tb_dados_compra (cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra) values %s`
	
	if err := prepare(query, dados); err != nil {
		return err
	}

	return nil

}

//Essa função é necessária para preparar a string com as colunas e valores para persistir os dados no banco
func prepare(s string, dados []*dadosCompra) error {

	db, err := db.Conectar();
	if err != nil {
		return err
	}
	defer db.Close()

	bf := bytes.Buffer{}
	values := make([]interface{}, 0, len(dados))
	for i, d := range dados {
		values = append(values,	
							d.Cpf,
							d.Private,
							d.Incompleto,
							d.DataUltimaCompra,
							d.TicketMedio,
							d.TicketUltimaCompra,
							d.LojaMaisFrequente,
							d.LojaUltimaCompra,)

		//Quantidade de colunas que serão persistidas no banco
		numFields := 8
		n := i * numFields

		bf.WriteString("(")
		for j := 0; j < numFields; j++ {
			bf.WriteString("$")
			bf.WriteString(strconv.Itoa(n + j + 1))
			bf.WriteString(", ")
		}
		bf.Truncate(bf.Len()-2)
		bf.WriteString("), ")
	}
	bf.Truncate(bf.Len()-2)

	stmt := fmt.Sprintf(s, bf.String())

	_, err = db.Exec(stmt, values...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}