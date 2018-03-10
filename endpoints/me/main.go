package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/satori/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/nullseed/lesshomeless-backend/helpers"
	"github.com/nullseed/lesshomeless-backend/models"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	userService, err := helpers.CreateUserService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	userID, ok := request.Headers["Authorization"]
	if !ok {
		u, err := userService.CreateUser()
		if err != nil {
			log.Printf("error creating user: %v\n", err)
			return helpers.CreateInternalServerErrorResponse()
		}

		return createMeResponse(u.Id)
	}

	u, err := userService.GetUser(userID)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return helpers.CreateUnauthorizedResponse()
	}

	offerService, err := helpers.CreateOfferService()
	if err != nil {
		log.Printf("error creating offer service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	giving, err := offerService.GetOffersById(u.Offers)
	if err != nil {
		log.Printf("error getting user's offers: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	o := meResponse{
		Giving: giving,
		Reserved: &models.Offer{
			Id:        uuid.NewV4().String(),
			Name:      string("Hey come reserve this"),
			CreatedOn: time.Now(),
			CreatedBy: uuid.NewV4().String(),
			Location: models.Location{
				Lat:  52.948956,
				Long: -1.150940,
			},
			Reservation: &models.Reservation{
				ReservedBy:   uuid.NewV4().String(),
				ReservedOn:   time.Now(),
				Acknowledged: false,
			},
		},
	}

	data, err := json.Marshal(o)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

type meResponse struct {
	Giving   []models.Offer `json:"giving"`
	Reserved *models.Offer  `json:"reserved"`
	UserID   string         `json:"userId,omitempty"`
}

func createMeResponse(userID string) (events.APIGatewayProxyResponse, error) {
	resp := meResponse{
		Giving:   nil,
		Reserved: nil,
		UserID:   userID,
	}

	data, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 500,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
