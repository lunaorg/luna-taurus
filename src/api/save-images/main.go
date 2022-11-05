package main

import (
	"context"
	"github.com/lunaorg/luna-taurus/src/libs/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

type Response struct {
	Resp string `json:"response"`
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := responses.NewResponses(ctx, "save-images-api")

	extraFields := map[string]interface{}{
		"received-event": event,
	}
	return resp.Success(extraFields)
}
