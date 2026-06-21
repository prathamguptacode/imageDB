package services

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func Upload(c fiber.Ctx) error {

	file, err := c.FormFile("upload")
	if err != nil {
		log.Println("Cannot upload image")
		return c.Status(400).JSON(fiber.Map{"message": "Something went wrong cannot upload file"})
	}

	allowedType := []string{"image/jpeg", "image/png"}
	errT := ValidateFileType(file, allowedType)
	if errT != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid file type"})
	}

	ext := filepath.Ext(file.Filename)
	safeName := uuid.New().String()
	fileName := safeName + ext

	errF := c.SaveFile(file, "upload/"+fileName)

	if errF != nil {
		log.Println("Cannot save image to disk")
		return c.JSON(fiber.Map{"message": "Something went wrong cannot save to disk"})
	}

	return c.JSON(fiber.Map{"message": "File uploaded successfully", "fileName": fileName})

}
