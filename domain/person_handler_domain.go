package domain

import "github.com/gofiber/fiber/v2"

type PersonHandler interface {
	GetPersonWithPersonID(c *fiber.Ctx) error
	InsertPerson(c *fiber.Ctx) error
	UpdatePerson(c *fiber.Ctx) error
	DeletePerson(c *fiber.Ctx) error
}
