package oss

import (
	"context"
	"fmt"
	"lightyear/core/global"
	"lightyear/pkg/oss"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitOssClient() {

	var accountId = global.Config.Oss.AccountId
	var accessKeyId = global.Config.Oss.AccessKeyId
	var accessKeySecret = global.Config.Oss.AccessKeySecret

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
			SigningRegion: "auto",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	s3client := s3.NewFromConfig(cfg)

	global.OssClient = &oss.OssClient{Client: s3client}

}
