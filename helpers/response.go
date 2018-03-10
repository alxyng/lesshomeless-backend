package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type errorResponse struct {
	Error string `json:"error"`
}

func CreateErrorResponse() (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: "Internal server error",
	}

	data, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: http.StatusInternalServerError,
	}, nil
}

func CreateUnauthorizedResponse() (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: "Unauthorized",
	}

	data, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: http.StatusUnauthorized,
	}, nil
}

func CreateBadRequestResponse() (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: "Bad request",
	}

	data, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: http.StatusBadRequest,
	}, nil
}

func CreateNotFoundResponse() (events.APIGatewayProxyResponse, error) {
	resp := errorResponse{
		Error: "Not found",
	}

	data, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: http.StatusNotFound,
	}, nil
}
