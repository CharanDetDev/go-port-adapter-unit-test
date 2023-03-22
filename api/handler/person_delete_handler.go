package handler

import (
	"strconv"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/caller"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/message"

	"github.com/gofiber/fiber/v2"
)

func (personHandler *personHandler) DeletePerson(c *fiber.Ctx) error {

	personId, err := strconv.Atoi(c.Params("personId"))
	if err != nil {
		return caller.BadRequestValidation(c, message.InvalidParam, "PersinID error.", err.Error())
	}

	if err := personHandler.PersonService.DeletePerson(personId); err != nil {
		return caller.InternalServerError(c, message.InvalidParam, err.Error())
	}

	return caller.Success(c, message.Success, message.Get(c, message.Success), nil)
}
