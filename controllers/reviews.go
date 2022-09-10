package controllers

import (
	"github.com/ZaphCode/go-fiber-ps/database"
	"github.com/ZaphCode/go-fiber-ps/models"
	"github.com/gofiber/fiber/v2"
)

func CreateReviewController(ctx *fiber.Ctx) error {
	body := new(models.Review)

	//* Get the request body
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Bad request",
		})
	}

	var product models.Product

	//* Check if the product exists
	if err := database.DB.Where("ID = ?", body.ProductID).First(&product).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That product does'nt exists",
		})
	}

	//* Create review
	newReview := models.Review{
		ProductID: body.ProductID,
		Comment:   body.Comment,
		Stars:     body.Stars,
	}

	if err := database.DB.Create(&newReview).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Error creating reviw",
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "review created",
		Data:    newReview,
	})
}

func GetReviewsController(ctx *fiber.Ctx) error {
	var reviews []models.Review

	//* Find all reviews and handle posible error
	if err := database.DB.Find(&reviews).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Searching reviews error",
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "All reviews",
		Data:    reviews,
	})
}

func SoftDeleteReviewController(ctx *fiber.Ctx) error {
	review_id := ctx.Params("id") // Parameter

	var review models.Review

	//* Find review
	if err := database.DB.Where("ID = ?", review_id).First(&review).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That review does'nt exists: " + err.Error(),
		})
	}

	//* Delete review
	if err := database.DB.Delete(&review).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Deleting error: " + err.Error(),
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Deleted",
		Data:    review,
	})
}

func PermanentDeleteReviewController(ctx *fiber.Ctx) error {
	review_id := ctx.Params("id") // Parameter

	var review models.Review

	//* Find review
	if err := database.DB.Unscoped().Where("ID = ?", review_id).First(&review).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "That review does'nt exists: " + err.Error(),
		})
	}

	//* Delete review
	if err := database.DB.Unscoped().Delete(&review).Error; err != nil {
		return ctx.Status(400).JSON(Response{
			Status:  "fail",
			Message: "Deleting error: " + err.Error(),
		})
	}

	//* Success
	return ctx.Status(201).JSON(Response{
		Status:  "Success",
		Message: "Deleted",
		Data:    review,
	})
}
