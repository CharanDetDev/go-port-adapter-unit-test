package config

import (
	"reflect"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func TestConfigEnv(t *testing.T) {

	ConfigInit()

	// marshal and print logger
	logg.PrintloggerJsonMarshalIndentHasHeader("\t\t\t ***** Test GET Config *****", "Environment", Env)

	// marshal ด้วย logg.JsonMarshalIndent(data interface{}) และ print logger ด้วย logg.Printlogger(header, prefix string, data interface{})
	logg.Printlogger("\t\t\t ***** Test GET Config :: Environment *****", converse.JsonMarshalIndent(Env))
}

func TestConfigValueWithKey(t *testing.T) {

	ConfigInit()

	// Find Value By Key (กำหนดชื่อ feild ใน struct ไว้ให้ตรงกับคือ)
	key := "API_PORT"

	val := reflect.ValueOf(&Env).Elem()
	for i := 0; i < val.NumField(); i++ {
		if key == val.Type().Field(i).Name {

			// print logger โดยการรับค่ามาเป็น Variadic function
			logg.PrintloggerVariadicJsonMarshalIndentHasHeader("\t\t ***** Test GET :: VALUE-With-KEY :: Config *****", "Result", converse.ParseToString_KeyValue("Type", val.Type().Field(i).Type), converse.ParseToString_KeyValue("Key", val.Type().Field(i).Name), converse.ParseToString_KeyValue("Value", val.Field(i).Interface()))

		}
	}

}

func TestConfigKeyWithValue(t *testing.T) {

	ConfigInit()

	// Find Key By Value (Type ต้องตรงกับค่าของ Key ที่ต้องการค้นหา)(กำหนดชื่อ feild ใน struct ไว้ให้ตรงกับคือ)
	value := ":3001"

	val := reflect.ValueOf(&Env).Elem()
	for i := 0; i < val.NumField(); i++ {
		if value == val.Field(i).Interface() {

			// print logger โดยการรับค่ามาเป็น Variadic function
			logg.PrintloggerVariadicJsonMarshalIndentHasHeader(
				"\t\t ***** Test GET :: KEY-With-VALUE :: Config *****",
				"Result",
				converse.ParseToString_KeyValue("Type", val.Type().Field(i).Type),
				converse.ParseToString_KeyValue("Key", val.Type().Field(i).Name),
				converse.ParseToString_KeyValue("Value", val.Field(i).Interface()))

		}
	}

}

func TestConfigInit(t *testing.T) {
	// var ttt ConfigTesting
	tests := []struct {
		name string
		num  int
	}{
		{"json", 0},
		{"ymal", 0},
		{"env", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ConfigInit(tt.name, ttt)s
		})
	}
}
