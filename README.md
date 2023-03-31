# The Go Language Programming : Port & Adapter, Unit Testing

> ## **Package**
>
> _Fiber, Mockgen, GORM, MySQL, Redis Cache, Viper_

- github.com/gofiber/fiber/v2
- github.com/golang/mock/gomock
- github.com/stretchr/testify/assert
- github.com/golang/mock/mockgen@v1.6.0
- gorm.io/gorm
- gorm.io/driver/mysql
- gorm.io/gorm/logger
- github.com/go-sql-driver/mysql
- github.com/redis/go-redis/v9"
- github.com/spf13/viper

> ## **vscode Setting**

- Package Mock & Generate สำหรับสร้างการจำลองข้อมูล
  - Go 1.16+ :
    go install github.com/golang/mock/mockgen@v1.6.0
  - Go version < 1.16 :
    GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
- Package Mock & Compare results :
  - go get github.com/stretchr/testify/assert
- Monitor viewer On vscode :
  แก้ไขเพิ่มเติมที่ setting.json :

```json
{
  "go.coverOnSave": true,
  "go.coverOnSingleTest": true,
  "go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
  }
}
```

- เปิด Terminal ณ ตำแหน่ง Package ที่ต้องการสร้าง Function Mock สำหรับทดสอบ
- Run คำสั่ง mockgen --source=filename.go --destination=./mock/filename.go จะได้ mock folder และ file ที่ถูกสร้างขึ้นจากคำสั่ง mockgen

- หลักการ AAA ของการทำ Unit test
  - Arrage : คือการเตรียมข้อมูลสำหรับการทำ Unit test
  - Action : คือการเรียกใช้ Function ที่ต้องการทดสอบ
  - Asset : คือการตรวจสอบผลลัพธ์ ว่าเป็นไปตามที่คาดหวังหรือไม่

> ## **Example**

- _ตัวอย่าง Port & Adapter, Unit Testing อยู่ที่ package handler, service, repository_

- _Basic Unit Testing_

```golang

package main

import "testing"


func main() {
	defer database.ConnectionClose()

	hello := "Hello"
	fmt.Println(hello, HelloWorld(hello))

	app := fiber.New()
	router := route.NewRoute()
	router.InitRoute(app)

	app.Listen(config.Env.API_PORT)
}

func HelloWorld(hello string) string {
	if hello != "" {
		return "World"
	}

	return ""
}

func TestHelloWorld(t *testing.T) { //! Function TestHelloWorld จะถูกสร้างขึ้นมาอัตโนมัติ

	//* Arrage
	type args struct {
		hello string //! argument ที่ต้องการส่งไปทดสอบ
	}
	tests := []struct {
		name string //! ชื่อของ test case
		args args   //! explicit field ของ argument
		want string //! expected
	}{
		// TODO: Add test cases.
		{
			name: "Case_Hello_World",
			args: args{
				hello: "Hello",
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Action
			got := HelloWorld(tt.args.hello)

			//* Assert
			if assert.Equal(t, got, tt.want) {
				fmt.Printf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
