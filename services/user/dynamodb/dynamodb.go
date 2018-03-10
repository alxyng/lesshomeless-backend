package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
)

type DynamoDBUserService struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDBUserService(db *dynamodb.DynamoDB) DynamoDBUserService {
	return DynamoDBUserService{
		db: db,
	}
}

func (s DynamoDBUserService) CreateUser() (models.User, error) {
	u := models.User{
		Id: uuid.NewV4().String(),
	}

	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return models.User{}, errors.Wrap(err, "error marshalling user")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("users"),
	})
	if err != nil {
		return models.User{}, errors.Wrap(err, "error putting user")
	}

	return u, nil
}
