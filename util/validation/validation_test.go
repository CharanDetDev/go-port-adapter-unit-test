package validation

import (
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

type PersonRequestModel struct {
	PersonID  int    `json:"person_id" validate:"required"`
	LastName  string `json:"last_name" validate:"required,gte=2"`
	FirstName string `json:"first_name" validate:"required"`
	Address   string `json:"address"`
	City      string `json:"city"`
}

func TestValidate(t *testing.T) {
	output := PersonRequestModel{
		PersonID:  3,
		LastName:  "1qq",
		FirstName: "",
		Address:   "",
		City:      "",
	}
	tests := []struct {
		name    string
		args    PersonRequestModel
		wantErr bool
	}{
		{
			name:    "Validate",
			args:    output,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			if err := validate(tt.args); (err != nil) != tt.wantErr {
				logg.PrintloggerJsonMarshalIndentHasHeader("", "", err.Error())
			}
		})
	}

}
