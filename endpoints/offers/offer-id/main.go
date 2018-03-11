package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nullseed/lesshomeless-backend/helpers"
	"github.com/nullseed/lesshomeless-backend/models"
	"github.com/nullseed/lesshomeless-backend/services/offer"
	"github.com/nullseed/lesshomeless-backend/services/user"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	offerService, err := helpers.CreateOfferService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	userService, err := helpers.CreateUserService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	userID, ok := request.Headers["Authorization"]
	if !ok {
		return helpers.CreateUnauthorizedResponse()
	}

	user, err := userService.GetUser(userID)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if user == nil {
		return helpers.CreateUnauthorizedResponse()
	}

	if request.HTTPMethod == "DELETE" {
		return delete(offerService, userService, request.PathParameters["id"], user)
	} else {
		return get(request.PathParameters["id"], offerService)
	}
}

func delete(offerService offer.OfferService, userService user.UserService, offerID string, user *models.User) (events.APIGatewayProxyResponse, error) {
	err := offerService.CancelOffer(offerID)
	if err != nil {
		log.Printf("failed to cancel offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	user, err = userService.RemoveOfferFromUser(user, offerID)
	if err != nil {
		log.Printf("failed to remove offer from user: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func get(id string, svc offer.OfferService) (events.APIGatewayProxyResponse, error) {
	o, err := svc.GetOffer(id)
	if err != nil {
		log.Printf("error getting offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if o == nil {
		log.Printf("offer %v not found\n", id)
		return helpers.CreateNotFoundResponse()
	}

	data, err := json.Marshal(o)
	if err != nil {
		log.Printf("error marshalling offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
