package connector

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBConnector struct {
	Config config.DBConfig
}

func NewPostgresDBConnector(config config.DBConfig) *PostgresDBConnector {
	return &PostgresDBConnector{
		Config: config,
	}
}

func (p PostgresDBConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(p.getConnectionString())) //DB logger to be investigated.
}

func (p *PostgresDBConnector) getConnectionString() string {
	config := p.Config
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host(),
		config.Username(),
		config.Password(),
		config.Name(),
		config.Port(),
		config.SSLMode(),
		config.TimeZone())
}

var _ DBConnector = &PostgresDBConnector{}
