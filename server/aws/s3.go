package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func GetS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_AWS_ACCESS_KEY_ID"),
			os.Getenv("S3_AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		return nil, err
	}
	return s3.New(sess), nil
}
