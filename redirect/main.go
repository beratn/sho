package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Print("id: " + id)
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

	if l.Target == "" {
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}

	js, _ := json.Marshal(l)
	return events.APIGatewayProxyResponse{Body: string(js), StatusCode: 301, Headers: map[string]string{
		"Location": l.Target,
	}}, nil
}
