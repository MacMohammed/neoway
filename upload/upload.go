package upload

import (
	"fmt"
	"io"
	"neoway/model"
	"neoway/response"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//Tamanho máximo do arquivo a ser enviado
const MAX_UPLOAD_SIZE = 10 << 20

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Método not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		fmt.Println(err)
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Garantir que o arquivo seja fechado ao fim do processamento
	defer file.Close()

	path := "./upload/files"

	//Se o diretório já existir, então a pasta não será criada
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//Cria a pasta para onde os arquivos serão enviados
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			//Caso não seja possível criar a pasta 'uploads', devolve o status de erro do servidor
			response.Erro(w, http.StatusInternalServerError, err)
			return
		}
	}

	/*
		Foi necessário criar o nome e diretório do arquivo em uma variável para,
		posteriormente, passar como parâmentro da função que persiste os dados do arquivo no banco.
	*/
	fileName := fmt.Sprintf("./upload/files/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	destino, err := os.Create(fileName)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer destino.Close()

	_, err = io.Copy(destino, file)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	 
	if err = model.PersistData(fileName); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, "Arquivo carregado com sucesso.")
}