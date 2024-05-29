package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func JSON(req events.APIGatewayV2HTTPRequest, v any) error {
	var buf bytes.Buffer
	if req.IsBase64Encoded {
		body, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return err
		}
		buf.Write(body)
	} else {
		buf.WriteString(req.Body)
	}

	if err := json.NewDecoder(&buf).Decode(v); err != nil {
		return err
	}

	return nil
}
