package handler

import (
	"strconv"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/caller"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/message"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func (personHandler *personHandler) GetPersonWithPersonID(c *fiber.Ctx) error {

	personId, err := strconv.Atoi(c.Params("personId"))
	if err != nil {
		return caller.BadRequestValidation(c, message.InvalidParam, message.Get(c, message.InvalidParam), err.Error())
	}

	var resPerson model.Person
	err = personHandler.PersonService.GetPersonWithPersonID(personId, &resPerson)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return caller.Success(c, message.PersonNotFound, message.Get(c, message.PersonNotFound), nil)
		} else {
			return caller.InternalServerError(c, message.InvalidParam, err.Error())
		}
	}

	return caller.Success(c, message.Success, message.Get(c, message.Success), resPerson)
}
