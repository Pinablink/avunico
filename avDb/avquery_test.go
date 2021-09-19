package avDb_test

import (
	"avunico/avDb"
	"avunico/avmodels"
	"testing"
)

func TestCreateQueryInsert(t *testing.T) {

	feiraLivreTeste := avmodels.FeiraLivre{
		ID:         "901",
		LONGI:      "-1220",
		LATI:       "-23322",
		SETCENS:    "T",
		AREAP:      "A",
		CODDIST:    "C",
		DISTRITO:   "D",
		CODSUBPREF: "C",
		SUBPREFE:   "SS",
		REGIAO5:    "RE",
		REGIAO8:    "R8",
		NOME_FEIRA: "TESTE",
		REGISTRO:   "TR",
		LOGRADOURO: "TL",
		NUMERO:     "NN",
		BAIRRO:     "BB",
		REFERENCIA: "RTT",
	}

	strQuery := avDb.CreateQueryInsert("teste", feiraLivreTeste)

	if len(strQuery) == 0 {
		t.Error("Erro")
	}
}

func TestCreateQueryUpdate(t *testing.T) {

	feiraLivreTeste := avmodels.FeiraLivre{
		LONGI: "-1220",
	}

	strQuery := avDb.CreateQueryUpdate("3", "Teste", feiraLivreTeste)

	if len(strQuery) == 0 {
		t.Error("Erro")
	}
}

func TestCreateQuerySelect(t *testing.T) {
	strQuery := avDb.CreateQuerySelect("DISTRITO", "PENHA")

	if len(strQuery) == 0 {
		t.Error("Erro")
	}
}

func TestCreateQueryDelete(t *testing.T) {
	strQuery := avDb.CreateQueryDelete("10")

	if len(strQuery) == 0 {
		t.Error("Erro")
	}
}
