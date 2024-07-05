package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSservice struct {
	client *s3.Client
}

func (awsS3 AWSservice) FileUpload(bucknetname, bucketkey, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("error while opening file", err)
	} else {
		defer file.Close()
	}
	_, err = awsS3.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucknetname),
		Key:    aws.String(bucketkey),
		Body:   file,
	})
	if err != nil {
		log.Println("error while uploading file", err)
	}
	return err
}

func main() {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("india"))
	if err != nil {
		log.Println("error while loading aws config")
	}
	awsservice := AWSservice{
		client: s3.NewFromConfig(config),
	}
	err = awsservice.FileUpload("bucket name", "bucket key", "file name")
	if err != nil {
		log.Println("error while uploading file in aws", err)
	} else {
		log.Println("successfully uploaded file in aws")

	}
}
