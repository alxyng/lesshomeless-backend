package dynamodb

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
)

const tableName = "offers"

type DynamoDBOfferService struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDBOfferService(db *dynamodb.DynamoDB) DynamoDBOfferService {
	return DynamoDBOfferService{
		db: db,
	}
}

func (s DynamoDBOfferService) GetOffer(id string) (*models.Offer, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	output, err := s.db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err, "error getting offer")
	}
	if output.Item == nil {
		return nil, nil
	}

	var o models.Offer
	dynamodbattribute.UnmarshalMap(output.Item, &o)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling item")
	}

	return &o, nil
}

func (s DynamoDBOfferService) CreateOffer(o models.Offer) (models.Offer, error) {
	o.Id = uuid.NewV4().String()
	o.CreatedOn = time.Now()

	item, err := dynamodbattribute.MarshalMap(o)
	if err != nil {
		return o, errors.Wrap(err, "error marshalling offer")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return o, errors.Wrap(err, "error putting offer")
	}

	return o, nil
}
