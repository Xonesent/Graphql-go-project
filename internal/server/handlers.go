package server

import (
	order_graphql "graphql/internal/order/order_delivery/graphql"
	order_http "graphql/internal/order/order_delivery/http"
	"graphql/internal/order/order_repository"
	"graphql/internal/order/order_usecase"
	product_graphql "graphql/internal/product/product_delivery/graphql"
	product_http "graphql/internal/product/product_delivery/http"
	product_mng_repo "graphql/internal/product/product_repository/mongo"
	product_psql_repo "graphql/internal/product/product_repository/postgres"
	"graphql/internal/product/product_usecase"
	user_graphql "graphql/internal/user/user_delivery/graphql"
	"graphql/internal/user/user_delivery/http"
	"graphql/internal/user/user_repository"
	"graphql/internal/user/user_usecase"
)

func (s *Server) MapHandlers() {
	userPsqlRepo := user_repository.NewUserPsqlRepository(s.cfg, s.psqlDB)
	productMngRepo := product_mng_repo.NewProductMngRepository(s.cfg, s.mngDB)
	productPsqlRepo := product_psql_repo.NewProductPsqlRepository(s.cfg, s.psqlDB)
	orderPsqlRepo := order_repository.NewOrderPsqlRepository(s.cfg, s.psqlDB)

	userUC := user_usecase.NewUserUseCase(s.cfg, userPsqlRepo)
	productUC := product_usecase.NewProductUseCase(s.cfg, productPsqlRepo, productMngRepo)
	orderUC := order_usecase.NewOrderUseCase(s.cfg, orderPsqlRepo)

	userHDL := user_http.NewUserHandler(s.cfg, userUC)
	productHDL := product_http.NewProductHandler(s.cfg, productUC)
	orderHDL := order_http.NewOrderHandler(s.cfg, orderUC)

	userResolver := user_graphql.NewUserResolver(s.cfg, userUC)
	productResolver := product_graphql.NewProductResolver(s.cfg, productUC)
	orderResolver := order_graphql.NewOrderResolver(s.cfg, orderUC)

	userGroup := s.fiberApp.Group("user")
	productGroup := s.fiberApp.Group("products")
	orderGroup := s.fiberApp.Group("order")
	ordersGroup := s.fiberApp.Group("orders")

	user_http.MapUserRoutes(userGroup, userHDL)
	product_http.MapProductRoutes(productGroup, productHDL)
	order_http.MapOrderRoutes(orderGroup, orderHDL)

	user_graphql.HandleUserGraphql(userGroup, userResolver)
	product_graphql.HandleProductGraphql(productGroup, productResolver)
	order_graphql.HandleOrderGraphql(ordersGroup, orderResolver)
}
