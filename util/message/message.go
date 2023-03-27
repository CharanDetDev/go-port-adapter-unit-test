package message

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leonelquinteros/gotext"
)

func Get(c *fiber.Ctx, code string) string {
	language := string(c.Request().Header.Peek("Content-Language"))
	return loadMessage(language, code)
}

func loadMessage(language string, code string) string {
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("this path :: ", path)

	GettextPath := "/Users/D.Charan/Documents/github/go-port-adapter/util/message/SelectLanguage/"
	gotext.Configure(GettextPath, strings.ToLower(language), "port_adaptor")
	message := gotext.Get(code)
	// fmt.Println("this message notfound :: ", message)
	if message == code {
		if message == "-0000004" {
			if strings.ToLower(language) == "th" {
				return "ไม่พบข้อมูลสมาชิก."
			} else {
				return "Member information not found."
			}
		} else {
			if strings.ToLower(language) == "th" {
				return fmt.Sprintf("[%s] เกิดข้อผิดพลาดขึ้นในระบบ.", code)
			} else {
				return fmt.Sprintf("[%s] An error in the system.", code)
			}
		}

	}
	return message
}
