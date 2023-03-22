package service

import (
	"reflect"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/cache"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"

	"github.com/CharanDetDev/go-port-adapter-unit-test/model"
	"github.com/CharanDetDev/go-port-adapter-unit-test/repository"
)

func init() {

	config.ConfigInitForTest()
	isDatabase := database.InitDatabase()
	isCache := cache.InitCache()

	if isDatabase && isCache {
		logg.Printlogger_Variadic("\t ***** Initail :: Configuration & Database & Redis :: SUCCESS **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient)
	} else {
		logg.Printlogger_Variadic("\t ***** Initail :: Configuration & Database & Redis :: ERROR **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient)
	}
}

func Test_personService_GetPersonWithPersonID(t *testing.T) {

	newPersonRepo := repository.NewPersonRepo()
	service := NewPersonService(newPersonRepo)

	type args struct {
		personId int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Person
		wantErr bool
	}{
		// happy case 1
		{
			name: "Test_personService_GetPersonWithPersonID",
			args: args{
				personId: 1,
			},
			wantErr: false,
		},
		// happy case 2
		{
			name: "Test_personService_GetPersonWithPersonID",
			args: args{
				personId: 2,
			},
			wantErr: false,
		},
		// fail
		{
			name: "Test_personService_GetPersonWithPersonID",
			args: args{
				personId: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetPersonWithPersonID(tt.args.personId)
			if err != nil {
				t.Logf("personService.GetPersonWithPersonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("personService.GetPersonWithPersonID() = %v, want %v", got, tt.want)
			}
		})
	}
}
