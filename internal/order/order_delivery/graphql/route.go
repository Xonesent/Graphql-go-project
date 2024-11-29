package order_graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
	"time"
)

func HandleOrderGraphql(group fiber.Router, resolver *OrderResolver) {
	h := handler.NewDefaultServer(NewExecutableSchema(newOrderConfig(resolver)))
	h.AddTransport(transport.SSE{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	group.All("/graphql", func(c *fiber.Ctx) error {
		wrapHandler(h.ServeHTTP)(c)
		return nil
	})
}

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

func newOrderConfig(resolver *OrderResolver) Config {
	c := Config{}
	c.Resolvers = resolver
	c.Directives.ValidateProdutId = ValidateProdutId

	return c
}
