package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/reward-rabieth/Authclk/api/handlers"
	"github.com/reward-rabieth/Authclk/api/middlewares"
	"github.com/reward-rabieth/Authclk/infrastucture/dataStores"
	"github.com/reward-rabieth/Authclk/infrastucture/identity"
	"github.com/reward-rabieth/Authclk/use_cases/productsuc"
	"github.com/reward-rabieth/Authclk/use_cases/usermgmtuc"
	"net/http"
)

func InitPublicRoute(c *echo.Echo) {
	c.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my demo Rest Api")
	})
	grp := c.Group("/api/v1")

	identityManager := identity.NewIdentityManger()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)
	grp.POST("/user", handlers.RegisterHandler(registerUseCase))

}

func InitProtectedRoute(c *echo.Echo) {
	grp := c.Group("/api/v1")

	productDataStore := dataStores.NewProductStore()
	createProductUseCase := productsuc.NewProductUseCase(productDataStore)
	grp.POST("/products", handlers.CreateProductHandler(createProductUseCase), middlewares.NewRequiresRealmRole("admin"))

	getProductsUseCase := productsuc.NewGetProductsUseCase(productDataStore)
	grp.GET("/products", handlers.GetProductsHandler(getProductsUseCase), middlewares.NewRequiresRealmRole("viewer"))
}
