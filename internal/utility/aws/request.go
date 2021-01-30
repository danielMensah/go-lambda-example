package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest

func Decode(req []byte, i interface{}) error {
	err := json.Unmarshal(req, &i)

	if err != nil {
		return err
	}

	return nil
}