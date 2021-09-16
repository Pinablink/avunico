package main

import (
	"avunico/avenvironment"
	"avunico/avlog"
	"fmt"
)

// Irá conter todos os parâmetros relacionados a base de dados
var mDbEnvironment *avenvironment.AvEnvDB
var mObLog *avlog.AvLog

func confRecursos() {
	mObLog = &avlog.AvLog{}
	mObLog.IniciarLogger()
	mObLog.Info().Printf("Recurso de Log Inicializado com sucesso!\n")
}

// Carrega e realiza a configuração de todos os parâmtros necessários a aplicação
func confParametros() {
	mDbEnvironment = &avenvironment.AvEnvDB{ObLogger: mObLog}
	param1, param2, merror := mDbEnvironment.CarregarParamEnv()

	if merror != nil {
		panic(merror)
	}

	fmt.Println("Passou do Ponto do código")
	fmt.Println(param1)
	fmt.Println(param2)
}

func main() {
	confRecursos()
	confParametros()
}
