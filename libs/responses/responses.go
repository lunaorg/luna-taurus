package responses

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Responses struct {
	functionName string
	extraFields  map[string]interface{}
}

func NewResponses(ctx context.Context, functionName string) (r *Responses) {
	return &Responses{
		functionName: functionName,
	}
}

func (r *Responses) Success(extraFields map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("%#v", Responses{
			functionName: r.functionName,
			extraFields:  extraFields,
		}),
		StatusCode: http.StatusOK,
	}, nil
}
