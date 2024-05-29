package response

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func JSON(status int, data interface{}) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: buf.String(),
	}, nil
}

func Status(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
	}, nil
}

func Text(status int, text string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: text,
	}, nil
}
