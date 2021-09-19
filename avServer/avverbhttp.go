package avServer

import (
	"avunico/avDb"
	"avunico/avlog"
	"avunico/avmodels"
	"encoding/json"
	"net/http"
	"strings"
)

// Contrato para cada tipo de verbo vindo do http request
type AvVerb interface {
	RunRequest(log *avlog.AvLog, db *avDb.AvDb, w http.ResponseWriter, r *http.Request) interface{}
}

// Atende um pedido de Get
type AvGet struct {
}

// Atende um pedido de Post
type AvPost struct {
}

// Atende um pedido de Atualização
type AvPut struct {
}

// Atende um pedido de exclusão
type AvDelete struct {
}

// Realiza o processamento da consulta
func (avGet AvGet) RunRequest(log *avlog.AvLog, db *avDb.AvDb, w http.ResponseWriter, r *http.Request) interface{} {
	urlReq := r.URL.String()
	log.Info().Printf("Processamento de Solicitação GET - URL: [%s]", urlReq)
	dadosReq := strings.Split(urlReq, "/")
	var aResponse avmodels.AvResponse

	if len(dadosReq) == 6 {
		param := dadosReq[4]
		val := dadosReq[5]

		col, err := db.SelectFeira(param, val)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
		}

		if len(col) > 0 {
			w.WriteHeader(http.StatusOK)
			return avmodels.AvResponseConsulta{Cod: 200, Message: "OK", ListaFeiraLivre: col}
		} else if len(col) == 0 {
			w.WriteHeader(http.StatusOK)
			return avmodels.AvResponse{Cod: 200, Message: "Não foi encontrada nenhuma feira com essa solicitação"}
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
	}

	return aResponse
}

// Realiza o processamento da inclusão de uma feira livre
func (avGet AvPost) RunRequest(log *avlog.AvLog, db *avDb.AvDb, w http.ResponseWriter, r *http.Request) interface{} {
	urlReq := r.URL.String()
	log.Info().Printf("Processamento de Solicitação POST - URL: [%s]", urlReq)
	var aResponse avmodels.AvResponseFeira
	var mFeiraLivre avmodels.FeiraLivre

	json.NewDecoder(r.Body).Decode(&mFeiraLivre)
	err := db.InsertFeira(mFeiraLivre)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
	}

	w.WriteHeader(http.StatusOK)
	aResponse = avmodels.AvResponseFeira{Cod: 200, Message: "Dados Incluídos com sucesso", ObInc: mFeiraLivre}

	return aResponse
}

// Realiza o processamento da atualização do registro de uma feira
func (avGet AvPut) RunRequest(log *avlog.AvLog, db *avDb.AvDb, w http.ResponseWriter, r *http.Request) interface{} {
	urlReq := r.URL.String()
	log.Info().Printf("Processamento de Solicitação PUT - URL: [%s]", urlReq)
	dadosReq := strings.Split(urlReq, "/")
	var aResponse avmodels.AvResponse

	if len(dadosReq) == 5 {
		var mFeiraLivre avmodels.FeiraLivre

		valId := dadosReq[4]
		json.NewDecoder(r.Body).Decode(&mFeiraLivre)
		err := db.UpdateFeira(valId, mFeiraLivre)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
		}

		w.WriteHeader(http.StatusOK)
		aResponse = avmodels.AvResponse{Cod: 200, Message: "Dados da Feira Alterados com Sucesso"}

	} else {
		w.WriteHeader(http.StatusNotFound)
		return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
	}

	return aResponse
}

// Realiza a exclusão de um registro de feira livre
func (avGet AvDelete) RunRequest(log *avlog.AvLog, db *avDb.AvDb, w http.ResponseWriter, r *http.Request) interface{} {
	urlReq := r.URL.String()
	log.Info().Printf("Processamento de Solicitação DELETE - URL: [%s]", urlReq)
	dadosReq := strings.Split(urlReq, "/")
	var aResponse avmodels.AvResponse

	if len(dadosReq) == 5 {
		valId := dadosReq[4]

		err := db.DeleteFeira(valId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
		}

		w.WriteHeader(http.StatusOK)
		aResponse = avmodels.AvResponse{Cod: 200, Message: "Feira excluida com sucesso"}

	} else {
		w.WriteHeader(http.StatusNotFound)
		return avmodels.AvResponse{Cod: 404, Message: "Recurso não encontrado"}
	}

	return aResponse
}
