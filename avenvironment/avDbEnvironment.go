package avenvironment

import (
	"avunico/avlog"
	"errors"
	"os"
)

// Contém os dados de acesso ao banco de dados carregados do Ambiente Operacional
type AvEnvDB struct {
	ObLogger   *avlog.AvLog
	usuarioDB  string
	passwordDB string
}

// Obtém os dados de acesso ao banco de dados pelo registrados no ambiente operacional
func (avenvdb *AvEnvDB) CarregarParamEnv() (string, string, error) {

	avenvdb.obterDadosAmbiente()
	usuarioDBNil := (len(avenvdb.usuarioDB) == 0)
	passawordDBNil := (len(avenvdb.passwordDB) == 0)

	if usuarioDBNil && passawordDBNil {
		avenvdb.ObLogger.Error().Printf("Error na obtenção dos dados de conexão com a base configurado no ambiente. Usuário retornado %s - Senha Retornada %s\n", avenvdb.usuarioDB, avenvdb.passwordDB)
		return "", "", errors.New("Parâmetros de Acesso a base não encontrados")
	}

	return "Opa", "Olha nós aqui", nil

}

func (avenvdb *AvEnvDB) obterDadosAmbiente() {
	avenvdb.usuarioDB = os.Getenv("AVUNICO_USUARIO_DB")
	avenvdb.passwordDB = os.Getenv("AVUNICO_PASSWORD_DB")
}
