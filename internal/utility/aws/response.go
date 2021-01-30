package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response events.APIGatewayProxyResponse

func OkResponse(body string) Response {
	return Response{
		StatusCode: http.StatusOK,
		Body: body,
	}
}

func ErrorResponse(e *Error, message string) Response {
	log.WithError(e.Error).Error(message)

	return Response{
		StatusCode: e.StatusCode,
		Body: e.ToJson(),
	}
}

func JsonResponse(data interface{}) Response {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return ErrorResponse(InternalServerError.SetError(err), "marshalling response")
	}

	return OkResponse(string(jsonData))
}