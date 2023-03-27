package repository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/domain"
	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/stretchr/testify/assert"
)

func Init(t *testing.T) {

	config.ConfigInitForTest()

	if config.Env.SKIP_UNIT_TEST == "" {
		t.Skip("skipping integration test")
	} else {
		database.InitDatabase()
	}
}

func Test_personRepo_GetPersonWithPersonID(t *testing.T) {

	//* Arrage
	Init(t)
	defer database.ConnectionClose()

	newPersonRepo := NewPersonRepo()
	type args struct {
		personId int
	}
	tests := []struct {
		name           string
		repo           domain.PersonRepo
		args           args
		expectedPerson *model.Person
		expectedError  error
	}{
		// TODO: Add test cases.
		{
			name: "Case_Success",
			repo: newPersonRepo,
			args: args{
				personId: 1,
			},
			expectedPerson: &model.Person{},
			expectedError:  nil,
		},
		{
			name: "Case_Not_Found",
			repo: newPersonRepo,
			args: args{
				personId: 100,
			},
			expectedPerson: &model.Person{},
			expectedError:  errors.New("record not found"),
		},
		{
			name: "Case_Internal_Server_Error",
			repo: newPersonRepo,
			args: args{
				personId: 1,
			},
			expectedPerson: &model.Person{},
			expectedError:  errors.New("sql: database is closed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Act
			if tt.name == "Case_Internal_Server_Error" {
				database.ConnectionClose()
			}
			err := tt.repo.GetPersonWithPersonID(tt.args.personId, tt.expectedPerson)

			//* Assert
			if assert.Equal(t, tt.expectedError, err) {
				if err != nil {
					logg.PrintloggerJsonMarshalIndentHasHeader(fmt.Sprintf("\t\t ***** Uni test :: %v *****", tt.name), "", fmt.Sprintf("Expected Error = %s, Actual = %s", tt.expectedError, err))
				} else {
					logg.PrintloggerJsonMarshalIndentHasHeader(fmt.Sprintf("\t\t ***** Uni test :: %v *****", tt.name), "", "Expected Error = nil, Actual = nil")
					logg.Printlogger("\t Result :: ", converse.JsonMarshalIndent(tt.expectedPerson))
				}
			}
			logg.PrintloggerUnderLineSingle()
		})
	}
}
