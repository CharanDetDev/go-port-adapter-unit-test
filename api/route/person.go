package route

import "github.com/gofiber/fiber/v2"

func (route *route) personGroup(person fiber.Router) {

	person.Get("/:personId", route.PersonHandler.GetPersonWithPersonID)
	person.Post("", route.PersonHandler.InsertPerson)
	person.Put("", route.PersonHandler.UpdatePerson)
	person.Delete("/:personId", route.PersonHandler.DeletePerson)

}
