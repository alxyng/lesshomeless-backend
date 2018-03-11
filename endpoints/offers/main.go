package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

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

	return get(offerService, request.QueryStringParameters["lat"], request.QueryStringParameters["long"])
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
		log.Printf("error creating offer: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	u, err = userSvc.AssignOfferToUser(u, o.Id)
	if err != nil {
		log.Printf("error assigning offer to user: %v\n", err)
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

func get(svc offer.OfferService, lat string, long string) (events.APIGatewayProxyResponse, error) {
	offers, err := svc.GetAllOffers()
	if err != nil {
		log.Printf("error getting all offers: %v\n", err)
		return helpers.CreateInternalServerErrorResponse()
	}

	var offersWithDistance []OfferWithDistance

	for _, o := range offers {
		/*var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		}
		var netClient = &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
		url := "http://dev.virtualearth.net/REST/v1/Routes/Walking?wayPoint.1=" + lat + "," + long + "&wayPoint.2=" + fmt.Sprint(o.Location.Lat) + "," + fmt.Sprint(o.Location.Long) + "&key=AheZ4uLLEL_fjKgtvICY0JBnbwYgJMxQe-u_54Pht6VEn1ZiI684arsTbjTR3t9n"
		fmt.Printf("Yooo - %v\n", url)
		resp, err := netClient.Get(url)
		if err != nil {
			log.Printf("error getting maps: %v\n", err)
			return helpers.CreateInternalServerErrorResponse()
		}

		var travelData BingMapsResponse
		if err := json.NewDecoder(resp.Body).Decode(&travelData); err != nil {
			log.Printf("error decoding distance body: %v\n", err)
			return helpers.CreateInternalServerErrorResponse()
		}
		fmt.Printf("Hullo - %v\n", travelData)
		offersWithDistance = append(offersWithDistance, OfferWithDistance{
			Offer:    o,
			Distance: travelData.ResourceSets[0].Resources[0].TravelDistance,
		})
		fmt.Printf("Dope! - %v\n", offersWithDistance)*/
		seedNum, _ := strconv.Atoi(string(o.Id)[0])
		fmt.Printf("Cool - %v\n", seedNum)

		rand.Seed(seedNum)
		offersWithDistance = append(offersWithDistance, OfferWithDistance{
			Offer:    o,
			Distance: rand.Float32() * 4,
		})
	}
	fmt.Printf("Holla dolla")
	data, err := json.Marshal(offersWithDistance)
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

type OfferWithDistance struct {
	Offer    models.Offer
	Distance float32
}

type BingMapsResponse struct {
	ResourceSets []ResourceSet `json:"resourceSets"`
}

type ResourceSet struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	TravelDistance float32 `json:"travelDistance"`
}
