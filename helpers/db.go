package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

func GetDynamoDBHandle() (*dynamodb.DynamoDB, error) {
	session, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://db:8000"),
		Region:   aws.String("eu-west-2"),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating aws session")
	}

	return dynamodb.New(session), nil
}
