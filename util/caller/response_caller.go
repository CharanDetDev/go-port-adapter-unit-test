package caller

import (
	"encoding/json"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/message"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/model"
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, code string, msg string, data interface{}) (err error) {
	return Response(c, fiber.StatusOK, &model.Response{
		Code:       code,
		Data:       data,
		HTTPStatus: fiber.StatusOK,
		Message:    msg,
	})
}

func BadRequest(c *fiber.Ctx, code string, err interface{}) error {
	return Response(c, fiber.StatusBadRequest, &model.ErrorResponseModel{
		Code:       code,
		HTTPStatus: fiber.StatusBadRequest,
		Message:    message.Get(c, code),
	})
}

func BadRequestValidation(c *fiber.Ctx, code string, errMsg string, err interface{}) error {
	return Response(c, fiber.StatusBadRequest, &model.ErrorResponseModel{
		Code:       code,
		HTTPStatus: fiber.StatusBadRequest,
		Message:    message.Get(c, code),
	})
}

func Unauthorized(c *fiber.Ctx, code string, err interface{}) error {
	return Response(c, fiber.StatusUnauthorized, &model.ErrorResponseModel{
		Code:       code,
		HTTPStatus: fiber.StatusUnauthorized,
		Message:    message.Get(c, code),
	})
}

func InternalServerError(c *fiber.Ctx, code string, err interface{}) error {
	return Response(c, fiber.StatusInternalServerError, &model.ErrorResponseModel{
		Code:       code,
		HTTPStatus: fiber.StatusInternalServerError,
		Message:    message.Get(c, code),
	})
}

func Response(c *fiber.Ctx, httpCode int, data interface{}) (err error) {
	js, _ := json.Marshal(data)
	c.Write(js)
	c.Status(httpCode)
	return nil
}
