
# avunico
# ESSA APLICAÇÃO, ASSIM COMO ESSA DOCUMENTAÇÃO, ESTÃO EM DESENVOLVIMENTO. TANTO O CÓDIGO FONTE QUANTO A DOC SERÃO EVOLUÍDOS E REVISADOS 
Microserviço que disponbiliza acesso a dados referentes a Localização de Feiras Livres do Município de São Paulo.

## Configurando alguns recursos
### Base de Dados

O Banco de Dados utilizado aqui é o MariaDB e as informações pertinentes a tabela que irá conter os dados para o funcionamento do microserviço. Foi extraído do arquivo http://www.prefeitura.sp.gov.br/cidade/secretarias/upload/chamadas/feiras_livres_1429113213.zip, disponibilizado pela Prefeitura de São Paulo.

#### Criação da Base de Dados

```
CREATE DATABASE `avunico` 
```
#### Criação da Tabela

👉  ***Importante:***
As colunas NUMERO, BAIRRO E REFERENCIA não estão com o mesmo tamanho de varchar em relação ao que esta definido no dicionário. Para realização da carga do CSV para evitar o erro de ***Data truncated for column*** eu defini um novo tamanho
____

```
CREATE TABLE `feiraslivres` (
	`ID` VARCHAR(8) NOT NULL COMMENT 'Número de identificação do estabelecimento georreferenciado por SMDU/Deinfo,SMDU/Deinfo',
        `LONGI` VARCHAR(10) NOT NULL COMMENT 'Longitude da localização do estabelecimento no território do Município, conforme MDC',
	`LATI` VARCHAR(10) NOT NULL COMMENT 'Latitude da localização do estabelecimento no território do Município, conforme MDC',
	`SETCENS` VARCHAR(15)  COMMENT 'Setor censitário,Setor censitário conforme IBGE',
	`AREAP` VARCHAR(13)  COMMENT 'Área de ponderação (agrupamento de setores censitários) conforme IBGE 2010',
	`CODDIST` VARCHAR(9) COMMENT 'Código do Distrito Municipal conforme IBGE',
	`DISTRITO`VARCHAR(18) COMMENT 'Nome do Distrito Municipal',
	`CODSUBPREF` VARCHAR(2) COMMENT 'Código de cada uma das 31 Subprefeituras (2003 a 2012)',
	`SUBPREFE` VARCHAR(25) COMMENT 'Nome da Subprefeitura (31 de 2003 até 2012)',
	`REGIAO5` VARCHAR(6)  COMMENT 'Região conforme divisão do Município em 5 áreas',
	`REGIAO8` VARCHAR(7)  COMMENT 'Região conforme divisão do Município em 8 áreas',
	`NOME_FEIRA` VARCHAR(30)  COMMENT 'Denominação da feira livre atribuída pela Supervisão de Abastecimento',
	`REGISTRO` VARCHAR(6) COMMENT 'Número do registro da feira livre na PMSP',
	`LOGRADOURO` VARCHAR(34) COMMENT 'Nome do logradouro onde se localiza a feira livre',
	`NUMERO` VARCHAR(20) COMMENT 'Um número do logradouro onde se localiza a feira livre',
	`BAIRRO` VARCHAR(30)  COMMENT 'Bairro de localização da feira livre',	
	`REFERENCIA` VARCHAR(40)  COMMENT 'Ponto de referência da localização da feira livre',
	PRIMARY KEY (`ID`)
)
COMMENT='Contém os os dados das feiras livres do Município de São Paulo'
COLLATE='latin1_swedish_ci'
```
#### Carga do CSV para a tabela na feiraslivres da base avunico

👉  ***Importante:***
No meu ambiente eu usei esse dir C:/CargaFeiraLivre/ para manter o arquivo DEINFO_AB_FEIRASLIVRES_2014.csv, conteúdo do feiras_livres_1429113213.zip. Essa configuração é da escolha do usuário.
___

```
LOAD DATA LOCAL INFILE 'C:/CargaFeiraLivre/DEINFO_AB_FEIRASLIVRES_2014.csv'
INTO TABLE feiraslivres
FIELDS TERMINATED BY ','
    ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 LINES
(ID,LONGI,LATI,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREFE,REGIAO5,REGIAO8,NOME_FEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA)
```

#### Variaveis de Ambiente

Os dados necessários para acesso ao banco de dados na minha configuração são mantidos como Variaveis de Ambiente do Sistema Operacional. Acredito que essa seja uma boa prática. No entanto fica a critério do usuário essa configuração.
