package controllers

import (
	"errors"
	"fiber-shop/internal/models"
	"fiber-shop/internal/services"
	"fiber-shop/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts returns all products
// @Description Returns all products
// @Error 500 {string} string "Internal server error"
// @Success 200 {object} models.Product
// @Router /products [get]
func GetAllProducts(c *fiber.Ctx) error {

	products, err := services.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(products)
}

// GetProductById returns product by ID
// @Description Returns a single product by ID
// @Param id path int true "Product ID"
// @Error 400 {string} string "Invalid product ID"
// @Error 404 {string} string "Product not found"
// @Error 500 {string} string "Internal server error"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	if !utils.IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	product, err := services.GetProductById(id)

	if err != nil {
		if err == utils.ErrorProductNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(product)
}

// PostProduct adds a new product to the database
// @Description Adds a new product to the database
// @Param product body models.Product true "Product information"
// @Success 201 {object} response "Product created successfully"
// @Error 400 {string} string "Name and price are required fields"
// @Error 500 {string} string "Product could not be added to the database"
// @Router /products [post]
func PostProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if product.Name == "" || product.Prices == nil { // TODO:validation
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name and price is required",
		})
	}

	p, err := services.AddProduct(&product)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Product can not be added",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Product created successfully",
		"product_id": p.Id,
	})
}

func GetProductPage(c *fiber.Ctx) error {
	//TODO:return page of products by page index
	//google pagination
	return errors.New("implement me")
}
