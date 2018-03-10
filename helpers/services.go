package helpers

import (
	"github.com/nullseed/lesshomeless-backend/services/offer"
	odb "github.com/nullseed/lesshomeless-backend/services/offer/dynamodb"
	"github.com/nullseed/lesshomeless-backend/services/user"
	udb "github.com/nullseed/lesshomeless-backend/services/user/dynamodb"
	"github.com/pkg/errors"
)

func CreateUserService() (user.UserService, error) {
	var service user.UserService

	db, err := GetDynamoDBHandle()
	if err != nil {
		return service, errors.Wrap(err, "error getting db handle")
	}

	return udb.NewDynamoDBUserService(db), nil
}

func CreateOfferService() (offer.OfferService, error) {
	var service offer.OfferService

	db, err := GetDynamoDBHandle()
	if err != nil {
		return service, errors.Wrap(err, "error getting db handle")
	}

	return odb.NewDynamoDBOfferService(db), nil
}
