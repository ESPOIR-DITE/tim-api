package env

import (
	"github.com/ESPOIR-DITE/tim-api/config/env"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
)

type TimApiEnvServiceConfiguration struct {
	env.EnvServiceConfiguration
	TimApiEnvServiceSQSConfig  SQSConfig
	TimApiServiceS3Config      S3Config
	TimApiServiceS3VideoConfig S3VideoConfig
}

var _ tim_api.ServiceConfiguration = &TimApiEnvServiceConfiguration{}

func (t *TimApiEnvServiceConfiguration) S3VideoConfig() tim_api.S3VideoConfig {
	return t.TimApiServiceS3VideoConfig
}

func (t *TimApiEnvServiceConfiguration) S3Config() tim_api.S3Config {
	return &t.TimApiServiceS3Config
}

func (t *TimApiEnvServiceConfiguration) SqsConfig() tim_api.SQSConfig {
	return &t.TimApiEnvServiceSQSConfig
}
