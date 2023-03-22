package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/message"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validate(v interface{}) error {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("%v", err.Error())
			return errors.New(errMessage)
		}
	}
	return nil
}

func Struct(c *fiber.Ctx, v interface{}, code string) (string, error) {
	err := json.Unmarshal(c.Body(), &v)
	if err != nil {
		logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("Unmarshal ERROR :: %v ", err.Error()))
		return message.Get(c, code), err
	}

	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, lcErr := lineAndCharacter(string(c.Body()), int(jsonError.Offset))
		logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("ERROR Message :: %v | %v | %v ", fmt.Sprintf("failed with error: Cannot parse JSON schema due to a syntax error at line %d", line), fmt.Sprintf("character %d", character), fmt.Sprintf("json error %v", jsonError.Error())))
		if lcErr != nil {
			logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("Couldn't find the line and character position of the error due to error %v", lcErr))
		}

	}
	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, lcErr := lineAndCharacter(string(c.Body()), int(jsonError.Offset))
		logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("ERROR Message :: %v | %v | %v | %v | %v | %v ", fmt.Sprintf("failed with error: The JSON type '%v'", jsonError.Value), fmt.Sprintf("cannot be converted into the Go '%v'", jsonError.Type.Name()), fmt.Sprintf("type on struct '%s'", jsonError.Struct), fmt.Sprintf("field  '%v'", jsonError.Field), fmt.Sprintf("See input file line %d", line), fmt.Sprintf("character %d", character)))

		if lcErr != nil {
			logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("failed with error: Couldn't find the line and character position of the error due to error %v", lcErr))
		}

	}
	if err != nil {
		logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("failed with error: %v", err))
		return message.Get(c, code), err
	}

	err = validate(v)
	if err != nil {
		logg.Printlogger("********** Validation ERROR **********", "", fmt.Sprintf("failed with error: %v", err))
		return message.Get(c, code), err
	}
	return "", nil
}

func lineAndCharacter(input string, offset int) (line int, character int, err error) {
	lf := rune(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("couldn't find offset %d within the input", offset)
	}

	line = 1

	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	return line, character, nil
}
