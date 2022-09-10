package main

import (
	"log"
	"os"

	"github.com/ZaphCode/go-fiber-ps/controllers"
	"github.com/ZaphCode/go-fiber-ps/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	app := fiber.New()

	database.ConnectDB()
	database.MigrateDB()

	//* Routes
	{
		router := app.Group("/api/products")
		router.Post("/create", controllers.CreateProductController)
		router.Get("/all", controllers.GetProductsController)
		router.Get("/get/:id", controllers.GetProductController)
		router.Delete("/delete/:id", controllers.SoftDeleteProductController)
		router.Delete("/permadelete/:id", controllers.PermanentDeleteProductController)
	}
	{
		router := app.Group("/api/reviews")
		router.Post("/create", controllers.CreateReviewController)
		router.Get("/all", controllers.GetReviewsController)
		router.Delete("/delete/:id", controllers.SoftDeleteReviewController)
		router.Delete("/permadelete/:id", controllers.PermanentDeleteReviewController)
	}

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
