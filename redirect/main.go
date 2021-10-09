package redirect

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/beratn/sho/model"
)

var app model.App

func init() {
	app = model.App{}
	app.Initialize()
}

func main() {
	lambda.Start(redirect)
}

func redirect(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, found := request.PathParameters["id"]

	if !found {
		return events.APIGatewayProxyResponse{Body: "Invalid ID", StatusCode: 401}, nil
	}

	l := model.Link{}
	res := app.Redis.Get(id).Val()

	if res == "" {
		l.GetTargetById(id)
		res = l.Target
		l.SetCache()
	}

	js, _ := json.Marshal(l)
	return events.APIGatewayProxyResponse{Body: string(js), StatusCode: 200}, nil
}
