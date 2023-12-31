package api

import (
	"fmt"
	"net/http"
	"orderapp/api/db"

	"github.com/labstack/echo/v4"
)

func getProduct(c echo.Context) error {
	name := c.Param("name")
	product, err := db.Get(name)

	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Error finding product %s", name))
	}

	return c.JSON(http.StatusOK, product)
}

func getAllProducts(c echo.Context) error {
	products, err := db.GetAll()
	if err != nil {
		return c.String(http.StatusNotFound, "Error finding products")
	}

	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context) error {
	p := new(db.Product)

	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	product := db.ProductDTO{
		Name:  p.Name,
		Price: p.Price,
	}

	err := db.Create(product)

	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Could not insert product %s", p.Name))
	}

	return c.String(http.StatusOK, fmt.Sprintf("Product %s inserted", p.Name))
}
