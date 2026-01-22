package database

import (
	"go-api-rest/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Database representa a conexão com o banco de dados
type Database struct {
	DB *gorm.DB
}

// NewDatabase cria uma nova conexão com o banco de dados
func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		return nil, err
	}

	logger.Info("Conexão com banco de dados estabelecida com sucesso")
	return &Database{DB: db}, nil
}

// Close fecha a conexão com o banco de dados
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
