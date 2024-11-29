package product_graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	fiberConn "graphql/pkg/dependencies/fiber"
)

func HandleProductGraphql(group fiber.Router, resolver *ProductResolver) {
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))

	group.All("/graphql", fiberConn.AdaptHTTPHandler(h))
}
