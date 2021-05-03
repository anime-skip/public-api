package main

import (
	"context"

	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var chiLambda *adapter.ChiLambda

func init() {
	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}

	router := server.GetRoutes(orm)
	chiLambda = adapter.New(router)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
