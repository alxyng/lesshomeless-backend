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

	offer, err := offerService.GetOffer(request.PathParameters["id"])
	if err != nil {
		log.Printf("error getting offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	if offer == nil {
		log.Printf("offer %v not found\n", request.PathParameters["id"])
		return helpers.CreateNotFoundResponse()
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

	if request.HTTPMethod == "PATCH" {
		return patch(offerService, offer)
	} else if request.HTTPMethod == "DELETE" {
		return delete(offerService, userService, offer, user)
	} else {
		return post(offerService, userService, offer, user)
	}
}

func patch(offerService offer.OfferService, offer *models.Offer) (events.APIGatewayProxyResponse, error) {
	if offer.Reservation == nil {
		log.Printf("cannot acknowledge offer without reservation\n")
		return helpers.CreateBadRequestResponse()
	}

	offer, err := offerService.AcknowledgeReservation(*offer)
	if err != nil {
		log.Printf("failed to reserve offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	data, err := json.Marshal(offer)
	if err != nil {
		log.Printf("error marshalling offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func delete(offerService offer.OfferService, userService user.UserService, offer *models.Offer, user *models.User) (events.APIGatewayProxyResponse, error) {
	offer, err := offerService.CancelReservation(*offer)
	if err != nil {
		log.Printf("failed to cancel reservation: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	user, err = userService.RemoveReservationFromUser(user, offer.Id)
	if err != nil {
		log.Printf("failed to remove reservation from user: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	data, err := json.Marshal(offer)
	if err != nil {
		log.Printf("error marshalling offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func post(offerService offer.OfferService, userService user.UserService, offer *models.Offer, user *models.User) (events.APIGatewayProxyResponse, error) {
	offer, err := offerService.ReserveOffer(*offer, user.Id)
	if err != nil {
		log.Printf("failed to reserve offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	user, err = userService.AssignReservationToUser(user, offer.Id)
	if err != nil {
		log.Printf("failed to assign reservation to user: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	data, err := json.Marshal(offer)
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
