package avDb

import (
	"fmt"
	"reflect"
)

//
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

// UPDATE `feiraslivres` SET SETCENS = 'Z', AREAP = 'Z' WHERE ID = '881'

//
func CreateQueryUpdate(nomeCampoId string, nomeTabela string, q interface{}) string {

	var updateQuery string

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		var valID string
		refVal := reflect.ValueOf(q)
		qtAtributos := refVal.NumField()

		updateQuery = fmt.Sprintf("UPDATE %s SET ", nomeTabela)

		for i := 0; i < qtAtributos; i++ {
			var r reflect.Value = refVal.Field(i)

			if len(r.String()) > 0 {
				campo := refVal.Type().Field(i).Name

				if campo == nomeCampoId {
					valID = r.String()
				} else {
					updateQuery = fmt.Sprintf("%s %s = \"%s\", ", updateQuery, campo, r.String())
				}

			}

		}

		runes := []rune(updateQuery)
		updateQuery = string(runes[:(len(runes) - 2)])
		updateQuery = fmt.Sprintf("%s WHERE %s = \"%s\" ", updateQuery, nomeCampoId, valID)

	}

	return updateQuery
}
