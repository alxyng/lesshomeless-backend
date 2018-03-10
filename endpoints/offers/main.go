package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/satori/uuid"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	o := []models.Offer{
		models.Offer{
			Id:        uuid.NewV4(),
			Name:      string("Cool offer bro"),
			CreatedOn: time.Now(),
			CreatedBy: uuid.NewV4(),
			Location: models.Location{
				Lat:  52.948956,
				Long: -1.150940,
			},
		},
		models.Offer{
			Id:        uuid.NewV4(),
			Name:      string("Hey come reserve this"),
			CreatedOn: time.Now(),
			CreatedBy: uuid.NewV4(),
			Location: models.Location{
				Lat:  52.948956,
				Long: -1.150940,
			},
			Reservation: &models.Reservation{
				ReservedBy:   uuid.NewV4(),
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

func main() {
	lambda.Start(handleRequest)
}
