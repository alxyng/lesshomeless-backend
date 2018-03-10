package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

func CreateInternalServerErrorResponse() (events.APIGatewayProxyResponse, error) {
	return createResponse(http.StatusInternalServerError, "Internal server error")
}

func CreateUnauthorizedResponse() (events.APIGatewayProxyResponse, error) {
	return createResponse(http.StatusUnauthorized, "Unauthorized")
}

func CreateBadRequestResponse() (events.APIGatewayProxyResponse, error) {
	return createResponse(http.StatusBadRequest, "Bad request")
}

func CreateNotFoundResponse() (events.APIGatewayProxyResponse, error) {
	return createResponse(http.StatusNotFound, "Not found")
}

type errorResponse struct {
	Error string `json:"error"`
}

func createResponse(statusCode int, msg string) (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: msg,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(data),
			StatusCode: http.StatusInternalServerError,
		}, errors.Wrap(err, "error marshalling response")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: statusCode,
	}, nil
}
