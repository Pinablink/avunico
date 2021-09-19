package avServer

import (
	"avunico/avDb"
	"avunico/avlog"
	"avunico/avmodels"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var mapProcessVerb map[string]AvVerb
var iObLog *avlog.AvLog
var iObDb *avDb.AvDb

//
type AvServer struct {
	obLog *avlog.AvLog
	obDb  *avDb.AvDb
}

//
func New(refObDb *avDb.AvDb, refOblog *avlog.AvLog) *AvServer {
	return &AvServer{obLog: refOblog,
		obDb: refObDb,
	}
}

// Inciliza o servidor
func (avserver *AvServer) InitServer() {

	iObLog = avserver.obLog
	iObDb = avserver.obDb
	fmt.Println("Inicializando o Processador de Verbos de Requisição")
	avserver.obLog.Info().Printf("Inicializando os Processador de Verbos de Requisição")

	mapProcessVerb = make(map[string]AvVerb)
	mapProcessVerb["GET"] = &AvGet{}
	mapProcessVerb["POST"] = &AvPost{}
	mapProcessVerb["PUT"] = &AvPut{}
	mapProcessVerb["DELETE"] = &AvDelete{}
	avserver.obLog.Info().Printf("Processador de Verbo Inicializado")

	http.HandleFunc("/", processRequest)
	error := http.ListenAndServe(":3000", nil)

	if error != nil {
		avserver.obLog.Error().Printf("Ocorreu erro na Inicialização do servidor")
		avserver.obLog.Error().Printf(error.Error())
		panic(error)
	}

}

// Valida se a requisição esta
func ValidarRequest(strUrlReq string) bool {

	dadosReq := strings.Split(strUrlReq, "/")

	if len(dadosReq) >= 4 {
		valApi := (dadosReq[1] == "api")
		valVersao := (dadosReq[2] == "v1")
		valEndPoint := (dadosReq[3] == "feiralivre")

		return (valApi && valVersao && valEndPoint)
	}

	return false
}

//
func processRequest(w http.ResponseWriter, r *http.Request) {
	urlReq := r.URL.String()
	w.Header().Set("Content-Type", "application/json")

	iObLog.Info().Printf("Solicitação de Serviço recebida")
	iObLog.Info().Printf("URL de solicitação %s", urlReq)

	if ValidarRequest(urlReq) {
		var processador AvVerb = mapProcessVerb[r.Method]
		json.NewEncoder(w).Encode(processador.RunRequest(iObLog, iObDb, w, r))

	} else {
		iObLog.Info().Printf("Recurso solicitado Inexistente")
		w.WriteHeader(http.StatusNotFound)
		avResponse := avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
		json.NewEncoder(w).Encode(avResponse)
	}

}
