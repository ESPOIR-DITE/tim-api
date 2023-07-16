package migration

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type PostgresMigrator struct {
	Config config.DBConfig
}

func NewPostgresMigrator(config config.DBConfig) *PostgresMigrator {
	return &PostgresMigrator{
		Config: config,
	}
}

func (p PostgresMigrator) NewMigrator() (*migrate.Migrate, error) {
	return migrate.New("file://"+p.Config.MigrationFolder(), p.connectionString())
}

func (p *PostgresMigrator) connectionString() string {
	config := p.Config
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Username(),
		config.Password(),
		config.Host(),
		config.Port(),
		config.Name(),
		config.SSLMode())
}

var _ Migrator = &PostgresMigrator{}
