package sqs

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"net"
	"strconv"
)

type SQSQueue struct {
	Config tim_api.SQSConfig
}

func NewSQSQueue(config tim_api.SQSConfig) *SQSQueue {
	return &SQSQueue{Config: config}
}

func (s *SQSQueue) NewSession() (*session.Session, error) {
	sqsEndpoint := fmt.Sprintf("%s://%s", s.Config.Protocol(), net.JoinHostPort(s.Config.Host(), strconv.Itoa(s.Config.Port())))

	result, err := session.NewSession(&aws.Config{
		Region:                        aws.String(s.Config.Region()),
		Endpoint:                      aws.String(sqsEndpoint),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}
