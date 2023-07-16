package env

import "github.com/ESPOIR-DITE/tim-api/config"

type EnvServiceConfiguration struct {
	EnvServiceAppConfig AppConfig
	EnvServiceDBConfig  DBConfig
}

var _ config.ServiceConfiguration = &EnvServiceConfiguration{}

func (e *EnvServiceConfiguration) AppConfig() config.AppConfig {
	return &e.EnvServiceAppConfig
}

func (e *EnvServiceConfiguration) DBConfig() config.DBConfig {
	return &e.EnvServiceDBConfig
}
