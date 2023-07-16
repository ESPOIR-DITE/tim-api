package tim_api

import "github.com/ESPOIR-DITE/tim-api/config"

type ServiceConfiguration interface {
	config.ServiceConfiguration
	SqsConfig() SQSConfig
	S3Config() S3Config
	S3VideoConfig() S3VideoConfig
}
