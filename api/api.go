package api

import (
	"fmt"
	"net/http"
	"orderapp/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Ping! :)")
}

func StartAPI() {
	//instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "joe" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	group := e.Group("/api/v1")

	// Routes
	e.GET("/ping", ping)
	group.GET("/product/:name", getProduct)
	group.GET("/products", getAllProducts)
	group.POST("/product", createProduct)
	// group.PUT("/product", updateProduct)
	// group.DELETE("/product", deleteProduct)

	// Start server
	port := fmt.Sprintf(":%d", config.Configuration.Server.Port)
	e.Logger.Fatal(e.Start(port))

}
