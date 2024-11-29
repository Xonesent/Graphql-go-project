package user_graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	fiberConn "graphql/pkg/dependencies/fiber"
)

func HandleUserGraphql(group fiber.Router, resolver *UserResolver) {
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))

	group.All("/graphql", fiberConn.AdaptHTTPHandler(h))
}
