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
// @Success 201                 object   response.Response{}
// @Router  /api/v1/file/upload [POST]
func (h FileHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
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
		}, file.Header["Content-Type"][0]) {
		return errors.BadRequestException(c, errors.INVALID_FILE_FORMAT)
	}

	if file.Size > config.File.MaxSize {
		return errors.PayloadTooLargeException(c)
	}

	stream, _ := file.Open()
	buffer := make([]byte, file.Size)
	stream.Read(buffer)

	// f, _ := os.Create("./" + file.Filename)
	// f.Write(buffer)

	// c.SaveFile(file, "./"+file.Filename)

	return response.WriteJSON(c, response.Response{})
}
