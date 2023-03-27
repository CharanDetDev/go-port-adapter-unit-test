package encryption

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func TestRandom32Char(t *testing.T) {

	random32char := RandomString(32)
	logg.Printlogger("\t\t ***** TEST random char32 *****", random32char)

}

func TestWriteNewEncryptionTo_Config_file(t *testing.T) {
	// fname := "./config.env"
	fname := "../config/config.env"
	// fname := "../config/config_dev.env"
	abs_fname, err := filepath.Abs(fname)
	if err != nil {
		logg.Printlogger("\t\t ***** TEST Reading Path File ERROR *****", err)
	}
	logg.Printlogger("\t\t ***** TEST Reading Path File SUCCESS *****", abs_fname)

	// file, err := os.OpenFile("./config.env", os.O_APPEND|os.O_WRONLY, 0644)
	file, err := os.OpenFile("../config/config.env", os.O_APPEND|os.O_WRONLY, 0644)
	// file, err := os.OpenFile("../config/config_dev.env", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logg.Printlogger("\t\t ***** TEST Open File ERROR *****", err)
	}
	defer file.Close()

	type args struct {
		key   string
		value string
	}

	tests := []args{
		{
			key:   "REDIS_DB_NUM",
			value: "0",
		},
	}

	for _, item := range tests {
		_ = EncryptParamsValue(item.value)
		// newLine := fmt.Sprintf("%v=%v", item.key, encrypt)
		// _, err = fmt.Fprintln(file, newLine)
		// if err != nil {
		// 	logg.Printlogger("\t\t ***** TEST Appending config to a file ERROR *****", "", err)
		// }

		// logg.Printlogger_Variadic_JsonMarshalIndent(
		// 	"\t\t ***** TEST Appending config to a file SUCCESS *****",
		// 	"Result For Test",
		// 	converse.ParseToString_KeyValue("Reading Path File", abs_fname),
		// 	converse.ParseToString_KeyValue("Open File", *file),
		// 	converse.ParseToString_KeyValue("Appending config", newLine),
		// )
	}

}

func TestEncryptParamsValue(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST Encryption",
			args: args{
				param: "kN6ulzxbp29cKMGTw5lMyLMzdz7jkqn5",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want = EncryptParamsValueConfig(tt.args.param, ""); tt.want == "" {
				logg.Printlogger("\t\t ***** EncryptParamsValue() ERROR *****", tt.want)
			}
			logg.Printlogger("\t\t ***** EncryptParamsValue() SUCCESS *****", tt.want)
		})
	}
}

func TestDecryptParamsValue(t *testing.T) {
	type args struct {
		DecryptString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST Decryption",
			args: args{
				DecryptString: "6fa8cb23e1f7be8d59839945c7012fa54dacfc57ea006380cc03341107",
			},
			want: "1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecryptParamsValue(tt.args.DecryptString)
			if got != tt.want {
				logg.PrintloggerVariadic("\t\t ***** DecryptParamsValue() ERROR *****", tt.want, got)
			}
			logg.PrintloggerVariadic("\t\t ***** DecryptParamsValue() SUCCESS *****", tt.want, got)
		})
	}
}
