package avDb

import (
	"avunico/avlog"
	"avunico/avmodels"
	"errors"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//
type AvDb struct {
	db         *sql.DB
	hostPortDb string
	usuarioDb  string
	passDb     string
	nomeDb     string
	obLog      *avlog.AvLog
}

// Inicializa o objeto de Base de Dados
func New(refStrHostPort string, refStrUsuario string, refStrPass string, refNome string, refObLog *avlog.AvLog) *AvDb {
	return &AvDb{
		hostPortDb: refStrHostPort,
		usuarioDb:  refStrUsuario,
		passDb:     refStrPass,
		nomeDb:     refNome,
		obLog:      refObLog,
	}
}

// Abre a conexão com a base de dados e realiza um teste de conexão
func (avdb *AvDb) Conn() error {
	var err error
	strDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", avdb.usuarioDb, avdb.passDb, avdb.hostPortDb, avdb.nomeDb)

	avdb.obLog.Info().Printf("Dados Utilizados para a criação do DSN. Usuário: %s - Password: %s - Host Port: %s - Nome da Base: %s", avdb.usuarioDb, avdb.passDb, avdb.hostPortDb, avdb.nomeDb)
	avdb.db, err = sql.Open("mysql", strDsn)

	// if there is an error opening the connection, handle it
	if err != nil {
		avdb.obLog.Error().Printf("Erro na Abertura de Conexão com a base de dados.")
		return err
	}

	avdb.obLog.Info().Printf("Abertura de Conexão OK")
	avdb.obLog.Info().Printf("Realizar teste de comunicação")
	err = avdb.db.Ping()

	if err != nil {
		avdb.obLog.Error().Printf("Erro na Comunicação com a Base de Dados")
		return err
	}

	avdb.obLog.Info().Printf("Comunicação realizada com sucesso")

	return nil
}

// Seleciona uma Feira Livre de acordo com a consulta informada
func (avdb *AvDb) SelectFeira(consulta string, valor string) (avmodels.FeiraLivre, error) {
	var feiraLivre avmodels.FeiraLivre

	if len(consulta) == 0 {
		avdb.obLog.Warning().Printf("Parâmetro de Consulta não foi informado")
		return feiraLivre, errors.New("Parâmetro de Consulta não informado na Solicitação de Feira")
	}

	if len(valor) == 0 {
		avdb.obLog.Warning().Printf("Valor do parâmetro %s para realizar a consulta de uma feira não foi infomado", consulta)
		return feiraLivre, errors.New("Valor do Parâmetro " + consulta + " para a consulta uma feira não foi informado")
	}

	strSelect := fmt.Sprintf("SELECT * FROM feiraslivres WHERE %s = ?", consulta)

	err := avdb.db.QueryRow(strSelect, valor).Scan(&feiraLivre.ID, &feiraLivre.LONGI, &feiraLivre.LATI,
		&feiraLivre.SETCENS, &feiraLivre.AREAP, &feiraLivre.CODDIST, &feiraLivre.DISTRITO,
		&feiraLivre.CODSUBPREF, &feiraLivre.SUBPREFE, &feiraLivre.REGIAO5, &feiraLivre.REGIAO8,
		&feiraLivre.NOME_FEIRA, &feiraLivre.REGISTRO, &feiraLivre.LOGRADOURO, &feiraLivre.NUMERO,
		&feiraLivre.BAIRRO, &feiraLivre.REFERENCIA)

	if err != nil {
		avdb.obLog.Error().Printf("Ocorreu erro na obtenção dos dados")
		avdb.obLog.Error().Printf(err.Error())
		return feiraLivre, err
	}

	return feiraLivre, nil
}

// Executa a Inserção de uma nova feira livre no cadastro
func (avdb *AvDb) InsertFeira(feiraLivre avmodels.FeiraLivre) error {
	strInsert := CreateQueryInsert("feiraslivres", feiraLivre)
	avdb.obLog.Info().Printf("Input de Feira Livre na Base Solicitado")
	avdb.obLog.Info().Printf("Query de Insert : %s", strInsert)

	insert, err := avdb.db.Query(strInsert)

	if err != nil {
		avdb.obLog.Error().Printf("Ocorreu um erro no processo de inclusão de registro")
		avdb.obLog.Error().Printf(err.Error())
		return err
	}

	insert.Close()

	return nil
}

//
func (avdb *AvDb) UpdateFeira(feiraLivre avmodels.FeiraLivre) error {
	strUpdate := CreateQueryUpdate("ID", "feiraslivres", feiraLivre)

	fmt.Println(strUpdate)
	return nil
}
