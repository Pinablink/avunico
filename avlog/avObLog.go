package avlog

import (
	"log"
	"os"
	"time"
)

var dataDoLog string

//
type AvLog struct {
	iFile         *os.File
	nomeLog       string
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
}

func obterDataCorrente() string {
	const formatDate = "20060102"
	var time1 = time.Now()
	time1 = time1.AddDate(0, 0, +1)
	return time1.Format(formatDate)
}

//
func New() *AvLog {
	return &AvLog{}
}

func (avlog *AvLog) IniciarLogger() {

	dataDoLog = obterDataCorrente()
	avlog.nomeLog = "log" + dataDoLog + ".txt"
	avlog.abrirLogAtual()
}

// Abrir arquivo de log atual
func (avlog *AvLog) abrirLogAtual() {
	var mErr error
	avlog.iFile, mErr = os.OpenFile(avlog.nomeLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if mErr != nil {
		panic(mErr)
	}

	avlog.warningLogger = log.New(avlog.iFile, "WARNING: ", log.LstdFlags|log.Lshortfile)
	avlog.infoLogger = log.New(avlog.iFile, "INFO: ", log.LstdFlags|log.Lshortfile)
	avlog.errorLogger = log.New(avlog.iFile, "ERROR: ", log.LstdFlags|log.Lshortfile)

}

// Fecha o Arquivo de Log Corrente.
func (avlog *AvLog) fecharLogAtual() {
	error := avlog.iFile.Close()

	if error != nil {
		panic("Erro sistêmico no controle de logs")
	}
}

// Retorna o Logger tipo Info
func (avlog *AvLog) Info() *log.Logger {

	data := obterDataCorrente()

	if dataDoLog != data {
		avlog.fecharLogAtual()
		dataDoLog = data
		avlog.nomeLog = "log" + dataDoLog + ".txt"
		avlog.abrirLogAtual()
	}

	return avlog.infoLogger
}

// Retorna o Logger tipo Warning
func (avlog *AvLog) Warning() *log.Logger {

	data := obterDataCorrente()

	if dataDoLog != data {
		avlog.fecharLogAtual()
		dataDoLog = data
		avlog.nomeLog = "log" + dataDoLog + ".txt"
		avlog.abrirLogAtual()
	}

	return avlog.warningLogger
}

// Retorna o Logger tipo Error
func (avlog *AvLog) Error() *log.Logger {

	data := obterDataCorrente()

	if dataDoLog != data {
		avlog.fecharLogAtual()
		dataDoLog = data
		avlog.nomeLog = "log" + dataDoLog + ".txt"
		avlog.abrirLogAtual()
	}

	return avlog.errorLogger
}
