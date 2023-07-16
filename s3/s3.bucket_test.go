package s3

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api"
	"github.com/ESPOIR-DITE/tim-api/config/tim_api/env"
	"os"
	"testing"
)

var S3Config tim_api.S3Config

func MyInit() *S3Bucket {
	S3Config := env.S3Config{}
	S3Config.S3Protocol = "http"
	S3Config.S3Host = "localhost"
	S3Config.S3Port = 4566
	session := NewS3Bucket("eu-west-2", S3Config)
	session.Init()
	return session
}
func TestCreateObject(t *testing.T) {
	file, err := os.ReadFile("ReadMe.mp4")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	session := MyInit()
	url, err := session.AddJsonPayload("tim-api-videos", "V-c38a7de0-075c-44c9-96eb-2f9808efa3d8", file)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(*url)
}
func TestS3Bucket_DownloadFile(t *testing.T) {

	session := MyInit()
	bytes, err := session.DownloadFile("tim-api-videos", "V-6d42e456-86b6-427a-a635-0606488cc722")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("ReadMe.mp4", bytes, 0666)
	if err != nil {
		fmt.Println(err)
	}

}
