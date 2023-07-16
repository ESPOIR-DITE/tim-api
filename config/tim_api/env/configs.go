package env

import "github.com/ESPOIR-DITE/tim-api/config/tim_api"

type SQSConfig struct {
	SQSProtocol  string `envconfig:"SQS_PROTOCOL" default:"http"`
	SQSHost      string `envconfig:"SQS_HOST"`
	SQSPort      int    `envconfig:"SQS_PORT"`
	SQSRegion    string `envconfig:"SQS_REGION"`
	SQSQueueName string `envconfig:"SQS_QUEUE_NAME"`
}

var _ tim_api.SQSConfig = &SQSConfig{}

func (s *SQSConfig) Host() string {
	return s.SQSHost
}

func (s *SQSConfig) Port() int {
	return s.SQSPort
}

func (s *SQSConfig) Protocol() string {
	return s.SQSProtocol
}

func (s *SQSConfig) Region() string {
	return s.SQSRegion
}

func (s *SQSConfig) QueueName() string {
	return s.SQSQueueName
}

type S3Config struct {
	S3Protocol   string `envconfig:"S3_PROTOCOL" default:"http"`
	S3Host       string `envconfig:"S3_HOST"`
	S3Region     string `envconfig:"S3_REGION"`
	S3BucketName string `envconfig:"S3_BUCKET_NAME"`
	S3Port       int    `envconfig:"S3_PORT"`
}

func (s S3Config) BucketName() string {
	return s.S3BucketName
}

func (s S3Config) Region() string {
	return s.S3Region
}

var _ tim_api.S3Config = &S3Config{}

func (s S3Config) Host() string {
	return s.S3Host
}

func (s S3Config) Port() int {
	return s.S3Port
}

func (s S3Config) Protocol() string {
	return s.S3Protocol
}

type S3VideoConfig struct {
	S3VideoProtocol   string `envconfig:"S3_VIDEO_PROTOCOL" default:"http"`
	S3VideoHost       string `envconfig:"S3_VIDEO_HOST"`
	S3VideoRegion     string `envconfig:"S3_VIDEO_REGION"`
	S3VideoBucketName string `envconfig:"S3_VIDEO_BUCKET_NAME"`
	S3VideoPort       int    `envconfig:"S3_VIDEO_PORT"`
}

func (s S3VideoConfig) Host() string {
	return s.S3VideoHost
}

func (s S3VideoConfig) Port() int {
	return s.S3VideoPort
}

func (s S3VideoConfig) Protocol() string {
	return s.S3VideoProtocol
}

func (s S3VideoConfig) BucketName() string {
	return s.S3VideoBucketName
}

func (s S3VideoConfig) Region() string {
	return s.S3VideoRegion
}

var _ tim_api.S3VideoConfig = S3VideoConfig{}
