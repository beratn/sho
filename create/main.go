package create

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/beratn/sho/model"
	"github.com/beratn/sho/utils"
	"strings"
)

func init() {
	a := model.App{}
	a.Initialize()
}

func main() {
	lambda.Start(createLink)
}

func createLink(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var l model.Link

	err := json.Unmarshal([]byte(request.Body), &l)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("Invalid Request"), StatusCode: 401}, nil
	}

	if !strings.HasPrefix(l.Target, "http") {
		l.Target = "http://" + l.Target
	}
	generatedId := utils.RandStringBytes(6)
	for model.CheckAddressIsExists(generatedId) {
		generatedId = utils.RandStringBytes(6)
	}
	l.Address = generatedId
	l.CreateLink()
	l.SetCache()

	js, _ := json.Marshal(l)
	return events.APIGatewayProxyResponse{Body: string(js), StatusCode: 200}, nil
}
