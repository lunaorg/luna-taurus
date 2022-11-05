package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lunaorg/luna-taurus/libs/responses"

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

	extraFields := map[string]interface{}{}
	err := json.Unmarshal([]byte(event.Body), &extraFields)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return resp.Success(extraFields)
}
