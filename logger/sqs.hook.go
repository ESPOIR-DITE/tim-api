package logger

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
)

type SQSHook struct {
	Session  *sqs.SQS
	QueueUrl *string
}

func (hook SQSHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook SQSHook) Fire(entry *logrus.Entry) error {
	// Send message to SQS
	sendMessageInput := sqs.SendMessageInput{}

	sendMessageInput.QueueUrl = hook.QueueUrl
	sendMessageInput.MessageBody = &entry.Message

	// We serialize data to JSON
	data, err := json.Marshal(&entry.Data)
	if err != nil {
		return fmt.Errorf("Failed to serialize log data into JSON")
	}

	sendMessageInput.MessageAttributes = map[string]*sqs.MessageAttributeValue{
		"Level": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(entry.Level.String()),
		},
		"Time": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(entry.Time.String()),
		},
		"Data": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(string(data)),
		},
	}
	_, err = hook.Session.SendMessage(&sendMessageInput)
	if err != nil {
		return err
	}

	return nil
}
