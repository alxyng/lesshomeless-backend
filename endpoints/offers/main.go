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

	userService, err := helpers.CreateUserService()
	if err != nil {
		log.Printf("error creating user service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	offerService, err := helpers.CreateOfferService()
	if err != nil {
		log.Printf("error creating offer service: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if request.HTTPMethod == "POST" {
		return post(request, userService, offerService)
	}

	return get(offerService)
}

func post(request events.APIGatewayProxyRequest, userSvc user.UserService, svc offer.OfferService) (events.APIGatewayProxyResponse, error) {
	var o models.Offer

	err := json.Unmarshal([]byte(request.Body), &o)
	if err != nil {
		log.Printf("error marshalling offer: %v\n", err)
		return helpers.CreateBadRequestResponse()
	}

	userID, ok := request.Headers["Authorization"]
	if !ok {
		return helpers.CreateUnauthorizedResponse()
	}

	u, err := userSvc.GetUser(userID)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}
	if u == nil {
		return helpers.CreateUnauthorizedResponse()
	}

	o.CreatedBy = u.Id

	o, err = svc.CreateOffer(o)
	if err != nil {
		log.Printf("error getting offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
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

func get(svc offer.OfferService) (events.APIGatewayProxyResponse, error) {
	offers, err := svc.GetAllOffers()
	if err != nil {
		log.Printf("error getting all offers: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	data, err := json.Marshal(offers)
	if err != nil {
		log.Printf("error marshalling offers: %v\n", err)
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
