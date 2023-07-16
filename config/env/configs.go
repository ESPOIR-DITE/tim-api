package env

import "github.com/ESPOIR-DITE/tim-api/config"

type AppConfig struct {
	AppPort      int    `envconfig:"HTTP_SERVER_PORT" required:"true"`
	AppAWSRegion string `envconfig:"AWS_REGION" default:"eu-west-1"`
	AppLogLevel  string `envconfig:"TIM_API_LOG_LEVEL" default:"INFO"`
}

var _ config.AppConfig = &AppConfig{}

func (config *AppConfig) Port() int {
	return config.AppPort
}

func (config *AppConfig) AWSRegion() string {
	return config.AppAWSRegion
}

func (config *AppConfig) LogLevel() string {
	return config.AppLogLevel
}

type DBConfig struct {
	DBUsername        string `envconfig:"DB_USERNAME" required:"true"`
	DBPassword        string `envconfig:"DB_PASSWORD" required:"true"`
	DBHost            string `envconfig:"DB_HOST" required:"true"`
	DBPort            int    `envconfig:"DB_PORT" required:"true"`
	DBName            string `envconfig:"DB_NAME" required:"true"`
	DBSSLMode         string `envconfig:"DB_SSLMODE" default:"disable"`
	DBTimeZone        string `envconfig:"DB_TIMEZONE" default:"Europe/London"`
	DBMigrationFolder string `envconfig:"DB_MIGRATION_FOLDER" default:"storage/databases/postgres/migrations"`
	DBLogLevel        string `envconfig:"DB_LOG_LEVEL" default:"silent"`
}

var _ config.DBConfig = &DBConfig{}

func (config *DBConfig) Port() int {
	return config.DBPort
}

func (config *DBConfig) Host() string {
	return config.DBHost
}

func (config *DBConfig) Username() string {
	return config.DBUsername

}

func (config *DBConfig) Password() string {
	return config.DBPassword
}

func (config *DBConfig) Name() string {
	return config.DBName
}

func (config *DBConfig) SSLMode() string {
	return config.DBSSLMode
}

func (config *DBConfig) TimeZone() string {
	return config.DBTimeZone
}

func (config *DBConfig) MigrationFolder() string {
	return config.DBMigrationFolder
}

func (config *DBConfig) LogLevel() string {
	return config.DBLogLevel
}
