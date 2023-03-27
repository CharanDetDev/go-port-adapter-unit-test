package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/domain"
	mock_domain "github.com/CharanDetDev/go-port-adapter-unit-test/domain/mock"
	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func mockService_GetPersonWithPersonID(t *testing.T, mockService *mock_domain.MockPersonService, expected error) {
	mockService.EXPECT().GetPersonWithPersonID(gomock.Any(), gomock.Any()).Return(expected)
}

func Test_personHandler_GetPersonWithPersonID(t *testing.T) {

	//* Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newFiber := fiber.New()

	newMockPersonService := mock_domain.NewMockPersonService(ctrl)
	newPersonHandler := NewPersonHandler(newMockPersonService)

	type args struct {
		requestPersonID string
	}
	tests := []struct {
		name               string
		app                *fiber.App
		personService      *mock_domain.MockPersonService
		personHandler      domain.PersonHandler
		args               args
		expectedPerson     *model.Person
		expectedStatusCode int
		expectedError      error
	}{
		// TODO: Add test cases.
		{
			name:          "Case_Success",
			app:           newFiber,
			personService: newMockPersonService,
			personHandler: newPersonHandler,
			args: args{
				requestPersonID: "1",
			},
			expectedPerson:     &model.Person{},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      nil,
		},
		{
			name:          "Case_BadRequest_InvalidParam",
			app:           newFiber,
			personService: newMockPersonService,
			personHandler: newPersonHandler,
			args: args{
				requestPersonID: "A",
			},
			expectedPerson:     &model.Person{},
			expectedStatusCode: fiber.StatusBadRequest,
			expectedError:      nil,
		},
		{
			name:          "Case_Not_Found",
			app:           newFiber,
			personService: newMockPersonService,
			personHandler: newPersonHandler,
			args: args{
				requestPersonID: "1000",
			},
			expectedPerson:     &model.Person{},
			expectedStatusCode: fiber.StatusOK,
			expectedError:      gorm.ErrRecordNotFound,
		},
		{
			name:          "Case_Internal_Server_Error",
			app:           newFiber,
			personService: newMockPersonService,
			personHandler: newPersonHandler,
			args: args{
				requestPersonID: "1",
			},
			expectedPerson:     &model.Person{},
			expectedStatusCode: fiber.StatusInternalServerError,
			expectedError:      errors.New("sql: database is closed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Act
			if tt.name != "Case_BadRequest_InvalidParam" {
				mockService_GetPersonWithPersonID(t, tt.personService, tt.expectedError)
			}

			req := httptest.NewRequest("GET", fmt.Sprintf("/%v", tt.args.requestPersonID), nil)
			tt.app.Get("/:personId?", tt.personHandler.GetPersonWithPersonID)
			res, _ := tt.app.Test(req)
			defer res.Body.Close()

			//* Assert
			if assert.Equal(t, tt.expectedStatusCode, res.StatusCode) {
				var responseBody *model.Person
				body, _ := io.ReadAll(res.Body)
				json.Unmarshal(body, &responseBody)
				logg.PrintloggerJsonMarshalIndentHasHeader(fmt.Sprintf("\t\t ***** Uni test :: %v *****", tt.name), "", fmt.Sprintf("Expected Status code = %v, Actual = %v", tt.expectedStatusCode, res.StatusCode))
			}
			logg.PrintloggerUnderLineSingle()
		})
	}
}
