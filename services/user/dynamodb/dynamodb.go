package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
)

const tableName = "lhl-users"

type DynamoDBUserService struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDBUserService(db *dynamodb.DynamoDB) DynamoDBUserService {
	return DynamoDBUserService{
		db: db,
	}
}

func (s DynamoDBUserService) CreateUser() (*models.User, error) {
	u := models.User{
		Id: uuid.NewV4().String(),
	}

	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling user")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error putting user")
	}

	return &u, nil
}

func (s DynamoDBUserService) GetUser(userID string) (*models.User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(userID),
			},
		},
		TableName: aws.String(tableName),
	}

	output, err := s.db.GetItem(input)
	if err != nil {
		return nil, errors.Wrap(err, "error getting user")
	}
	if output.Item == nil {
		return nil, nil
	}

	var u models.User
	dynamodbattribute.UnmarshalMap(output.Item, &u)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling item")
	}

	return &u, nil
}

func (s DynamoDBUserService) AssignOfferToUser(u *models.User, offerId string) (*models.User, error) {
	u.Offers = append(u.Offers, offerId)

	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling user")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error putting user")
	}

	return u, nil
}

func (s DynamoDBUserService) AssignReservationToUser(user *models.User, offerID string) (*models.User, error) {
	user.Reserved = append(user.Reserved, offerID)

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling user")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error putting user")
	}

	return user, nil
}

func (s DynamoDBUserService) RemoveReservationFromUser(user *models.User, offerID string) (*models.User, error) {
	for p, v := range user.Reserved {
		if v == offerID {
			user.Reserved[p] = user.Reserved[len(user.Reserved)-1]
			user.Reserved = user.Reserved[:len(user.Reserved)-1]
			break
		}
	}

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling user")
	}

	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error putting user")
	}

	return user, nil
}
