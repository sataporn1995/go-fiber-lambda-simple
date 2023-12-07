package main

import (
	// "log"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	app := fiber.New(fiber.Config{Prefork: false})

	app.Use(cors.New())

	// apiPrefix := app.Group("/api", logger.New())
	apiPrefix := app.Group("/api")

	apiPrefix.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"method": "GET"})
	})

	apiPrefix.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"method": "POST"})
	})

	apiPrefix.Put("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"method": "PUT"})
	})

	apiPrefix.Patch("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"method": "PATCH"})
	})

	apiPrefix.Delete("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"method": "DELETE"})
	})

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
	// log.Fatal(app.Listen(":3000"))
}
