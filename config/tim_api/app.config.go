package tim_api

type ClientConfig interface {
	Host() string
	Port() int
	Protocol() string
}
type SQSConfig interface {
	ClientConfig
	Region() string
	QueueName() string
}
type S3Configs interface {
	ClientConfig
	BucketName() string
	Region() string
}
type S3Config interface {
	S3Configs
}

type S3VideoConfig interface {
	S3Configs
}
