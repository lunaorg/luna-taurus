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

func NewResponses(ctx context.Context, functionName string) *Responses {
	return &Responses{
		functionName: functionName,
	}
}

func (r *Responses) Success(extraFields map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	r.extraFields = extraFields

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%#v", r),
		StatusCode: http.StatusOK,
	}, nil
}
