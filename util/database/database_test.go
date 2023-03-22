package database

import (
	"reflect"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

type Person struct {
	PersonID  int    `gorm:"column:PersonID"`
	LastName  string `gorm:"column:LastName"`
	FirstName string `gorm:"column:FirstName"`
	Address   string `gorm:"column:Address"`
	City      string `gorm:"column:City"`
}

func (m *Person) TableName() string {
	return "Persons"
}

func TestDatabaseGORM_Get_Person_With_PersonID(t *testing.T) {
	tests := []struct {
		name             string
		wantRowsAffected int64
		getByPersonId    int
		person           Person
	}{
		{
			name:             "Get Person with PersonID = 1",
			wantRowsAffected: 1,
			getByPersonId:    1,
		},
		{
			name:             "Get Person with PersonID = 2",
			wantRowsAffected: 1,
			getByPersonId:    2,
		},
	}

	config.ConfigInitForTest()
	InitDatabase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := Conn.First(&tt.person, tt.getByPersonId)
			if result.Error != nil {
				t.Errorf("GORM Get Person :: RowsAffected = %v, error = %v, want %v", result.RowsAffected, result.Error, 1)
			}

			if !reflect.DeepEqual(result.RowsAffected, tt.wantRowsAffected) {
				t.Errorf("GORM Get Person :: RowsAffected = %v, want %v", result.RowsAffected, tt.wantRowsAffected)
			}

			logg.Printlogger_JsonMarshalIndent(
				"\t\t TestDatabaseGORM_Get_Person_With_PersonID",
				"GET Person",
				tt.person)

		})
	}
}
