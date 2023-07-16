package config

type ServiceConfiguration interface {
	AppConfig() AppConfig
	DBConfig() DBConfig
}
