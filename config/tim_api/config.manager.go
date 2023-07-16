package tim_api

type ConfigManager interface {
	Load() (ServiceConfiguration, error)
}
