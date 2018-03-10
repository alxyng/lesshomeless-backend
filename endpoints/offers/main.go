package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nullseed/lesshomeless-backend/helpers"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/nullseed/lesshomeless-backend/services/offer"
	"github.com/satori/uuid"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	offerService, err := helpers.CreateOfferService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateErrorResponse()
	}

	if request.HTTPMethod == "POST" {
		return post(request.Body, offerService)
	} else {
		return get()
	}
}

func post(body string, svc offer.OfferService) (events.APIGatewayProxyResponse, error) {
	var o models.Offer
	err := json.Unmarshal([]byte(body), &o)

	if err == nil {
		fmt.Printf("Got offer: %#v\n", o)
	} else {
		fmt.Printf("Failed to deserialize offer %v\n", err)
	}

	o.CreatedBy = "01234567-1234-1234-1234-123456789012"

	o, err = svc.CreateOffer(o)
	if err != nil {
		log.Printf("error getting offer: %v\n", err)
		return helpers.CreateErrorResponse()
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

func get() (events.APIGatewayProxyResponse, error) {
	o := []models.Offer{
		models.Offer{
			Id:        uuid.NewV4().String(),
			Name:      string("Cool offer bro"),
			CreatedOn: time.Now(),
			CreatedBy: uuid.NewV4().String(),
			Location: models.Location{
				Lat:  52.948956,
				Long: -1.150940,
			},
		},
		models.Offer{
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

func main() {
	lambda.Start(handleRequest)
}
