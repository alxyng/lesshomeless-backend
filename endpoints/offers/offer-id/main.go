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

	if request.HTTPMethod == "DELETE" {
		return delete(request.PathParameters["id"])
	} else {
		return get(request.PathParameters["id"])
	}
}

func delete(id string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func get(id string) (events.APIGatewayProxyResponse, error) {
	o := models.Offer{
		Id:        id,
		Name:      string("Cool offer bro"),
		CreatedOn: time.Now(),
		CreatedBy: uuid.NewV4().String(),
		Location: models.Location{
			Lat:  52.948956,
			Long: -1.150940,
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
