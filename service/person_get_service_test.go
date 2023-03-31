package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/domain"
	mock_domain "github.com/CharanDetDev/go-port-adapter-unit-test/domain/mock"
	"github.com/CharanDetDev/go-port-adapter-unit-test/repository"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
)

func init() {
	config.ConfigInitForTest()
}

func mockRepo_GetPersonWithPersonID(t *testing.T, mockRepo *mock_domain.MockPersonRepo, expected error) {
	mockRepo.EXPECT().GetPersonWithPersonID(gomock.Any(), gomock.Any()).Return(expected)
}

func Test_personService_GetPersonWithPersonID(t *testing.T) {

	//* Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockPersonRepo := mock_domain.NewMockPersonRepo(ctrl)
	NewRedisCacheRepo := repository.NewRedisCacheRepo(database.RedisCaching)
	newPersonService := NewPersonService(newMockPersonRepo, NewRedisCacheRepo)

	type args struct {
		personId int
	}
	tests := []struct {
		name              string
		newMockPersonRepo *mock_domain.MockPersonRepo
		newPersonService  domain.PersonService
		args              args
		expectedPerson    *model.Person
		expectedErr       error
	}{
		{
			name:              "Case_Success",
			newMockPersonRepo: newMockPersonRepo,
			newPersonService:  newPersonService,
			args: args{
				personId: 1,
			},
			expectedPerson: &model.Person{},
			expectedErr:    nil,
		},
		{
			name:              "Case_Not_Found",
			newMockPersonRepo: newMockPersonRepo,
			newPersonService:  newPersonService,
			args: args{
				personId: 2,
			},
			expectedPerson: &model.Person{},
			expectedErr:    gorm.ErrRecordNotFound,
		},
		{
			name:              "Case_Internal_Server_Error",
			newMockPersonRepo: newMockPersonRepo,
			newPersonService:  newPersonService,
			args: args{
				personId: 2,
			},
			expectedPerson: &model.Person{},
			expectedErr:    errors.New("sql: database is closed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Act
			mockRepo_GetPersonWithPersonID(t, tt.newMockPersonRepo, tt.expectedErr)
			err := tt.newPersonService.GetPersonWithPersonID(tt.args.personId, tt.expectedPerson)

			//* Assert
			if assert.Equal(t, tt.expectedErr, err) {
				if err != nil {
					logg.PrintloggerJsonMarshalIndentHasHeader(fmt.Sprintf("\t\t ***** Uni test :: %v *****", tt.name), "", fmt.Sprintf("Expected Error = %v, Actual = %v", tt.expectedErr, err))
				} else {
					logg.PrintloggerJsonMarshalIndentHasHeader(fmt.Sprintf("\t\t ***** Uni test :: %v *****", tt.name), "", "Expected Error = nil, Actual = nil")
				}
			}

			logg.PrintloggerUnderLineSingle()
		})
	}
}
