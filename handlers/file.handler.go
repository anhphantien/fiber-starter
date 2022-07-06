package handlers

import (
	"fiber-starter/config"
	"fiber-starter/errors"
	"fiber-starter/response"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/slices"
)

type FileHandler struct{}

// @Tags    file
// @Summary Upload a file
// @Param   file                formData file false " "
// @Success 201                 object   response.Response{data=boolean}
// @Router  /api/v1/file/upload [POST]
func (h FileHandler) Upload(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		switch err {
		case fasthttp.ErrNoMultipartForm:
			return errors.BadRequestException(c, errors.FILE_NOT_FOUND)
		default:
			return errors.BadRequestException(c, err.Error())
		}
	}

	if !slices.Contains(
		[]string{
			config.File.ContentType.JPEG,
			config.File.ContentType.PNG,
		}, fileHeader.Header["Content-Type"][0]) {
		return errors.BadRequestException(c, errors.INVALID_FILE_FORMAT)
	}

	if fileHeader.Size > config.File.MaxSize {
		return errors.PayloadTooLargeException(c)
	}

	// file, _ := fileHeader.Open()
	// buffer := make([]byte, fileHeader.Size)
	// file.Read(buffer)
	// c.SaveFile(fileHeader, "./"+fileHeader.Filename)

	return response.WriteJSON(c, response.Response{
		Data: true,
	})
}
