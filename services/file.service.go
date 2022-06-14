package services

import (
	"fiber-starter/common"
	"fiber-starter/config"
	"fiber-starter/errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

type FileService struct{}

// @Tags    file
// @Summary Upload a file
// @Param   file            formData file false " "
// @Success 201             {object} common.HttpResponse{}
// @Router  /v1/file/upload [post]
func (s FileService) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		switch err.Error() {
		case "request has no multipart/form-data Content-Type":
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
		return errors.RequestEntityTooLargeException(c)
	}

	stream, _ := file.Open()
	buffer := make([]byte, file.Size)
	stream.Read(buffer)

	// f, _ := os.Create(fmt.Sprint("./", file.Filename))
	// f.Write(buffer)

	// c.SaveFile(file, fmt.Sprint("./", file.Filename))

	return c.JSON(common.HttpResponse{
		StatusCode: fiber.StatusCreated,
	})
}
