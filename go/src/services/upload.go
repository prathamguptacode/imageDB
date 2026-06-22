package services

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
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

	//catching images
	imgFile, errF := os.Open("upload/" + fileName)
	if errF != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	img, _, errI := image.Decode(imgFile)
	if errI != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	imgFile.Close()
	r := resize.Resize(800, 0, img, resize.Lanczos3)
	imgW, errw := os.Create("upload/" + safeName + "w_800.jpeg")
	if errw != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	errOw := jpeg.Encode(imgW, r, nil)
	if errOw != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	imgW.Close()
	r2 := resize.Resize(1600, 0, img, resize.Lanczos3)
	imgW2, errw2 := os.Create("upload/" + safeName + "w_1600.jpeg")
	if errw2 != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	errOw2 := jpeg.Encode(imgW2, r2, nil)
	if errOw2 != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	imgW2.Close()

	return c.JSON(fiber.Map{"message": "File uploaded successfully", "fileName": fileName})

}
