package database

import (
	"log"

	"github.com/matheusandre1/Lab-ci-go/config"
	"github.com/matheusandre1/Lab-ci-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao, configErr := config.BuildPostgresDSN()
	if configErr != nil {
		log.Panicf("Configuracao de ambiente invalida: %v", configErr)
	}

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
