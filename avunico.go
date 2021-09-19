package main

import (
	"avunico/avDb"
	"avunico/avServer"
	"avunico/avlog"
	"fmt"
	"os"
)

// Irá conter todos os parâmetros relacionados a base de dados

func confLog() *avlog.AvLog {
	var mObLog *avlog.AvLog = avlog.New()
	mObLog.IniciarLogger()
	return mObLog
}

//
func confDB(refObLog *avlog.AvLog) *avDb.AvDb {

	var obAvdb *avDb.AvDb

	strHostPort := os.Getenv("AVUNICO_HOST_PORT")
	strUsuario := os.Getenv("AVUNICO_USUARIO")
	strPass := os.Getenv("AVUNICO_PASSWORD")
	strDbName := os.Getenv("AVUNICO_DB_NAME")

	strHostPortOK := (len(strHostPort) > 0)
	strUsuarioOK := (len(strUsuario) > 0)
	strPassOK := (len(strPass) > 0)
	strDbNameOK := (len(strDbName) > 0)

	if strHostPortOK && strUsuarioOK && strPassOK && strDbNameOK {
		obAvdb = avDb.New(strHostPort, strUsuario, strPass, strDbName, refObLog)
	} else {
		refObLog.Error().Printf("Não foi possivel configurar a base de dados.")
		refObLog.Error().Printf("Parâmetros de configuração utilizados")
		refObLog.Error().Printf("Usuário: %s ; Senha: %s ; Host Port: %s ; Db : %s ;", strUsuario, strPass, strHostPort, strDbName)
		panic("Não foi possivel configurar a base de dados. Verifique os parâmetros de conexão")
	}

	return obAvdb
}

// Carrega e realiza a configuração de todos os parâmtros necessários a aplicação
func main() {
	var obLog *avlog.AvLog = confLog()
	avdbData := confDB(obLog)
	err := avdbData.Conn()

	if err != nil {
		panic(err)
	}

	fmt.Println("Inicializando o Servidor")
	var avServer *avServer.AvServer = avServer.New(avdbData, obLog)
	avServer.InitServer()
}
