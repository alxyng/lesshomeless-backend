package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nullseed/lesshomeless-backend/helpers"
	"github.com/nullseed/lesshomeless-backend/services/offer"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	offerService, err := helpers.CreateOfferService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateErrorResponse()
	}

	if request.HTTPMethod == "DELETE" {
		return delete(request.PathParameters["id"])
	} else {
		return get(request.PathParameters["id"], offerService)
	}
}

func delete(id string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func get(id string, svc offer.OfferService) (events.APIGatewayProxyResponse, error) {
	o, err := svc.GetOffer(id)
	if err != nil {
		log.Printf("error getting offer: %v\n", err)
		return helpers.CreateErrorResponse()
	}

	if o == nil {
		log.Printf("offer %v not found\n", id)
		return helpers.CreateNotFoundResponse()
	}

	data, err := json.Marshal(o)
	if err != nil {
		log.Printf("error marshalling offer: %v\n", err)
		return helpers.CreateErrorResponse()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
