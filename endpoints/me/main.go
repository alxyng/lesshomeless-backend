package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/satori/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/nullseed/lesshomeless-backend/models"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// userService, err := helpers.CreateUserService()
	// if err != nil {
	// 	log.Printf("error creating user service: %v\n", err)
	// 	return createErrorResponse()
	// }
	//
	// userService.CreateUser()

	o := models.Me{
		Giving: models.Offer{
			Id:        uuid.NewV4().String(),
			Name:      string("Cool offer bro"),
			CreatedOn: time.Now(),
			CreatedBy: uuid.NewV4().String(),
			Location: models.Location{
				Lat:  52.948956,
				Long: -1.150940,
			},
		},
		Reserved: models.Offer{
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

type errorResponse struct {
	Error string `json:"error"`
}

func createErrorResponse() (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: "Internal server error",
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
