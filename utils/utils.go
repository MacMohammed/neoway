package utils

import (
	"regexp"
	"strings"
)

//LimparCpfCnpj remove os caracteres inválidos ('.', '-', '/') e retorna uma string
func LimparCpfCnpj(cpf string) (string, error) {
	// s := strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(cpf), ".", "", -1), "-","",-1), "/","", -1)	
	// return s

	reg, err := regexp.Compile("[^0-9]+")
    if err != nil {
        return "", err
    }

	processedString := reg.ReplaceAllString(cpf, "0")

	return processedString, nil
}

func RemoveVirgula(v string) (string, error) {
	reg, err := regexp.Compile("[^0-9|,]+")
    if err != nil {
        return "", err
    }

	processedString := strings.Replace(reg.ReplaceAllString(v, "0"), ",", ".", -1)

	return processedString, nil
}

func RemoveNull(v string) (*string, error) {
	reg, err := regexp.Compile("[^0-9|-]+")
    if err != nil {
        return nil, err
    }

	processedString := reg.ReplaceAllString(v, "")

	if processedString == "" {
		return nil , nil
	}

	return &processedString, nil
}