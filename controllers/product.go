package controllers

import (
	"github.com/ZaphCode/go-fiber-ps/database"
	"github.com/ZaphCode/go-fiber-ps/models"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func CreateProductController(ctx *fiber.Ctx) error {
	body := new(models.Product)

	//* Get the request body
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Bad request",
		})
	}

	//* Create the product
	newProduct := models.Product{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	}

	if err := database.DB.Create(&newProduct).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Error creating product",
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Product created",
		Data:    newProduct,
	})
}

func GetProductsController(ctx *fiber.Ctx) error {
	var products []models.Product

	//* Find all products and populate the reviews
	if err := database.DB.Preload("Reviews").Find(&products).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Error searching products",
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "All products",
		Data:    &products,
	})
}

func GetProductController(ctx *fiber.Ctx) error {
	product_id := ctx.Params("id") // Parameter
	var product models.Product

	//* Find product
	if err := database.DB.Preload("Reviews").Where("ID = ?", product_id).First(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That product does'nt exists: " + err.Error(),
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Product " + product_id,
		Data:    product,
	})
}

func SoftDeleteProductController(ctx *fiber.Ctx) error {
	product_id := ctx.Params("id") // Parameter

	var product models.Product

	//* Find product
	if err := database.DB.Preload("Reviews").Where("ID = ?", product_id).First(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That product does'nt exists: " + err.Error(),
		})
	}

	//* Delete his reviews
	for _, review := range product.Reviews {
		if err := database.DB.Delete(&review).Error; err != nil {
			return ctx.Status(400).JSON(Response{
				Status:  "fail",
				Message: "Deleting review error: " + err.Error(),
			})
		}
	}

	//* Delete product
	if err := database.DB.Delete(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Deleting product error: " + err.Error(),
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Product and reviws deleted",
		Data:    product,
	})
}

func PermanentDeleteProductController(ctx *fiber.Ctx) error {
	product_id := ctx.Params("id") // Parameter

	var product models.Product

	//* Find product
	if err := database.DB.Unscoped().Preload("Reviews").Where("id = ?", product_id).First(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That product does'nt exists: " + err.Error(),
		})
	}

	//* Delete his reviews
	var reviews []models.Review

	if err := database.DB.Unscoped().Where("product_id = ?", product_id).Find(&reviews).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "WTF" + err.Error(),
		})
	}

	for _, review := range reviews {
		if err := database.DB.Unscoped().Delete(&review).Error; err != nil {
			return ctx.Status(400).JSON(Response{
				Status:  "fail",
				Message: "Deleting review error: " + err.Error(),
			})
		}
	}

	//* Delete product
	if err := database.DB.Unscoped().Delete(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Deleting product error: " + err.Error(),
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Product and reviws deleted permanently",
		Data:    product,
	})
}
