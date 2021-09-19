package avDb

import (
	"fmt"
	"reflect"
)

// Cria a Query de Inserção de uma nova feira livre na base de dados
func CreateQueryInsert(nomeTabela string, q interface{}) string {

	var insertQuery string

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		refVal := reflect.ValueOf(q)
		qtAtributos := refVal.NumField()

		insertQuery = fmt.Sprintf("INSERT INTO %s VALUES(", nomeTabela)

		for i := 0; i < qtAtributos; i++ {

			if i == 0 {
				insertQuery = fmt.Sprintf("%s \"%s\"", insertQuery, refVal.Field(i).String())
			} else {
				insertQuery = fmt.Sprintf("%s, \"%s\"", insertQuery, refVal.Field(i).String())
			}

		}

		insertQuery = fmt.Sprintf("%s)", insertQuery)
	}

	return insertQuery
}

// Cria a Query de atualização de uma feira livre na base de dados
func CreateQueryUpdate(valorID string, nomeTabela string, q interface{}) string {

	var updateQuery string

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		refVal := reflect.ValueOf(q)
		qtAtributos := refVal.NumField()

		updateQuery = fmt.Sprintf("UPDATE %s SET ", nomeTabela)

		for i := 0; i < qtAtributos; i++ {
			var r reflect.Value = refVal.Field(i)

			if len(r.String()) > 0 {
				campo := refVal.Type().Field(i).Name
				updateQuery = fmt.Sprintf("%s %s = \"%s\", ", updateQuery, campo, r.String())
			}

		}

		runes := []rune(updateQuery)
		updateQuery = string(runes[:(len(runes) - 2)])
		updateQuery = fmt.Sprintf("%s WHERE %s = \"%s\" ", updateQuery, "ID", valorID)

	}

	return updateQuery
}

// Cria a Query de seleção de uma feira livre
func CreateQuerySelect(strConsulta string, strValor string) string {
	return fmt.Sprintf("SELECT * FROM feiraslivres WHERE %s = \"%s\"", strConsulta, strValor)
}

// Cria a Query de exclusão de uma feira livre
func CreateQueryDelete(strId string) string {
	return fmt.Sprintf("DELETE FROM feiraslivres WHERE ID = \"%s\"", strId)
}
