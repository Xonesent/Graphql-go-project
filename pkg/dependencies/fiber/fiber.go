package fiberConn

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"graphql/pkg/constant"
	utils "graphql/pkg/utilities"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberConfig struct {
	Host string `envconfig:"FIBER_HOST" validate:"required"`
	Port string `envconfig:"FIBER_PORT" validate:"required"`
}

func NewFiberClient() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          FiberErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))
	//app.Use(helmet.New())
	//app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fullStackTrace := string(debug.Stack())
			log.Println("Recovered from panic: " + utils.LimitStackTrace(fullStackTrace, 1))
		},
	}))

	return app
}

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	setStatusCode(ctx, err)

	if utils.InStringSlice(constant.Host, constant.DevHosts) {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": nil,
	})
}

func setStatusCode(ctx *fiber.Ctx, err error) {
	statusCode := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		statusCode = e.Code
	}

	ctx.Status(statusCode)
}

func AdaptHTTPHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(c.Context())
		return nil
	}
}
