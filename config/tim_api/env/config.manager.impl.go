package env

import (
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type TimApiConfigurationManagerImpl struct{}

var _ tim_api.ConfigManager = &TimApiConfigurationManagerImpl{}

func NewTimApiConfigurationManagerImpl() *TimApiConfigurationManagerImpl {
	return &TimApiConfigurationManagerImpl{}
}

func (t TimApiConfigurationManagerImpl) Load() (config tim_api.ServiceConfiguration, err error) {
	config = new(TimApiEnvServiceConfiguration)
	if err = envconfig.Process("", config); err != nil {
		return
	}
	return
}
