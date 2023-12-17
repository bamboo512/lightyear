package oss

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type OssClient struct {
	Client *s3.Client
}

func (client *OssClient) UploadFile(bucketName string, key string, filePath string) error {

	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	// 设置一个超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	_, err = client.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   file,
	})

	if err != nil {
		return err
	}
	return nil
}

func (client *OssClient) UploadLargeObject(bucketName string, objectKey string, largeObject []byte) error {
	largeBuffer := bytes.NewReader(largeObject)
	var partMiBs int64 = 10
	uploader := manager.NewUploader(client.Client, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
	})
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   largeBuffer,
	})
	if err != nil {
		log.Printf("Couldn't upload large object to %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}

	return err
}

// DownloadFile to local path
func (client *OssClient) DownloadFile(bucketName string, key string, filePath string) error {

	result, err := client.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, key, err)
		return err
	}
	defer result.Body.Close()
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", filePath, err)
		return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", key, err)
	}
	_, err = file.Write(body)
	return err
}
