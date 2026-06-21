package services

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/HugoSmits86/nativewebp"
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
	wU, errWU := strconv.ParseUint(width, 10, 64)
	if errWU != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	ext := filepath.Ext(myFilePath)
	imgType := strings.Split(ext, ".")[1]
	var img image.Image
	var errG error
	if imgType == "jpg" || imgType == "jpeg" {
		img, errG = jpeg.Decode(file)
	}
	if imgType == "png" {
		img, errG = png.Decode(file)
	}
	if errG != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	r := resize.Resize(uint(wU), 0, img, resize.Lanczos3)
	var buf bytes.Buffer
	errW := nativewebp.Encode(&buf, r, nil)
	if errW != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}
	c.Set("Content-Type", "image/webp")
	return c.Send(buf.Bytes())

}
