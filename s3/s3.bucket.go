package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type S3Bucket struct {
	AwsRegion  string
	S3Config   tim_api.S3Configs
	client     *s3.Client
	s3Endpoint string
}

func NewS3Bucket(awsRegion string, s3Config tim_api.S3Configs) *S3Bucket {
	return &S3Bucket{
		AwsRegion: awsRegion,
		S3Config:  s3Config,
	}
}

func (b *S3Bucket) Init() error {
	if b.S3Config.Host() != "" {
		b.s3Endpoint = fmt.Sprintf("%s://%s", b.S3Config.Protocol(), net.JoinHostPort(b.S3Config.Host(), strconv.Itoa(b.S3Config.Port())))
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		{
			if strings.TrimSpace(b.s3Endpoint) != "" {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           b.s3Endpoint,
					SigningRegion: b.AwsRegion,
				}, nil

			}
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		}

	})

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(b.AwsRegion),
		awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return fmt.Errorf("AWS LoadDefaultConfig Error: %s", err.Error())
	}

	b.client = s3.NewFromConfig(cfg,
		func(o *s3.Options) {
			o.UsePathStyle = true
		})
	return nil
}

func (b *S3Bucket) AddJsonPayload(bucketName, key string, payload []byte) (*string, error) {
	if b.client == nil {
		return nil, fmt.Errorf("S3 bucket client is not yet initialized")
	}
	_, err := b.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   bytes.NewReader(payload),
	})

	if err != nil {
		return nil, fmt.Errorf("S3 error: %s", err.Error())
	}

	bucketURL := b.GetBucketURL(bucketName, key)
	return &bucketURL, nil
}

func (b *S3Bucket) GetBucketURL(bucketName, key string) string {
	if strings.TrimSpace(b.s3Endpoint) == "" {
		return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, b.AwsRegion, key)
	}
	return fmt.Sprintf("%s/%s/%s", b.s3Endpoint, bucketName, key)
}

// DownloadFile gets an object from a bucket and stores it in a local file.
func (b *S3Bucket) DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	result, err := b.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return nil, err
	}
	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body, nil
}
