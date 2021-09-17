package avmodels

//
type FeiraLivre struct {
	ID         string `json: "id"`
	LONGI      string `json: "long"`
	LATI       string `json: "lat"`
	SETCENS    string `json: "setcens"`
	AREAP      string `json: "areap"`
	CODDIST    string `json: "coddist"`
	DISTRITO   string `json: "distrito"`
	CODSUBPREF string `json: "codsubpref"`
	SUBPREFE   string `json: "subpref"`
	REGIAO5    string `json: "regiao05"`
	REGIAO8    string `json: "regiao08"`
	NOME_FEIRA string `json: "nome_feira"`
	REGISTRO   string `json: "registro"`
	LOGRADOURO string `json: "logradouro"`
	NUMERO     string `json: "numero"`
	BAIRRO     string `json: "bairro"`
	REFERENCIA string `json: "referencia"`
}
