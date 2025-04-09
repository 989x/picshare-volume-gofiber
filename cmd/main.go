// cmd/main.go

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// UploadHandler handles image upload to /static/images/
func UploadHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Image is required")
	}

	dirPath := "./static/images"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Println("Directory creation failed:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
	}

	dst := filepath.Join(dirPath, file.Filename)
	if err := c.SaveFile(file, dst); err != nil {
		log.Println("File save failed:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save image")
	}

	return c.JSON(fiber.Map{
		"message": "Image uploaded successfully",
		"path":    "/images/" + file.Filename,
	})
}

// HealthCheckHandler confirms the server is running
func HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API connected successfully",
	})
}

func main() {
	app := fiber.New()

	app.Post("/upload", UploadHandler)
	app.Get("/", HealthCheckHandler)

	log.Println("Server is running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
