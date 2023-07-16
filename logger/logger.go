package logger

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
	sqs2 "github.com/ESPOIR-DITE/tim-api/sqs"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func LogInit(config tim_api.ServiceConfiguration) error {

	Log.SetFormatter(&logrus.JSONFormatter{})

	session, err := sqs2.NewSQSQueue(config.SqsConfig()).NewSession()
	if err != nil {
		Log.Fatal("Failed to get sqs session : %w", err.Error())
		return err
	}

	client := sqs.New(session)

	sqsURL := fmt.Sprintf("%d/%s", client.Config.Endpoint, config.SqsConfig().QueueName())

	Log.AddHook(SQSHook{client, &sqsURL})

	Log.SetOutput(os.Stdout)
	return nil
}
