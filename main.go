package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/controller"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.POST("/", controller.AddOrder)
	r.GET("/", controller.GetOrders)
	r.GET("/:id", controller.GetOrder)
	ginLambda = ginadapter.New(r)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
