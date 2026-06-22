package services

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/nfnt/resize"
)

func Retrieve(c fiber.Ctx) error {

	fileReq := c.Params("file")
	myFilePath := "upload/" + fileReq

	file, errF := os.Open(myFilePath)
	if errF != nil {
		return c.Status(400).JSON(fiber.Map{"message": "File not found"})
	}
	defer file.Close()

	m := c.Queries()
	width := m["width"]
	if width == "" {
		return c.SendFile(myFilePath)
	}
	if width == "800" {
		ext := filepath.Ext(myFilePath)
		fileN := strings.Split(fileReq, ext)
		fileNameP := "upload/" + fileN[0] + "w_800.jpeg"
		return c.SendFile(fileNameP)
	}
	if width == "1600" {
		ext := filepath.Ext(myFilePath)
		fileN := strings.Split(fileReq, ext)
		fileNameP := "upload/" + fileN[0] + "w_1600.jpeg"
		return c.SendFile(fileNameP)
	}
	wU, errWU := strconv.ParseUint(width, 10, 64)
	if errWU != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid width"})
	}

	img, _, errG := image.Decode(file)
	if errG != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	r := resize.Resize(uint(wU), 0, img, resize.Lanczos3)
	var buf bytes.Buffer
	errW := jpeg.Encode(&buf, r, nil)
	if errW != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	c.Set("Content-Type", "image/jpeg")
	return c.Send(buf.Bytes())

}
