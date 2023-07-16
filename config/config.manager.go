package config

type ConfigManager interface {
	Load() (ServiceConfiguration, error)
}
