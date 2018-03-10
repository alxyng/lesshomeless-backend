package helpers

import (
	"github.com/nullseed/lesshomeless-backend/services/user"
	"github.com/nullseed/lesshomeless-backend/services/user/dynamodb"
	"github.com/pkg/errors"
)

func CreateUserService() (user.UserService, error) {
	var service user.UserService

	db, err := GetDynamoDBHandle()
	if err != nil {
		return service, errors.Wrap(err, "error getting db handle")
	}

	return dynamodb.NewDynamoDBUserService(db), nil
}
