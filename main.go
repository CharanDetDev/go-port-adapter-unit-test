package main

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/api/route"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/gofiber/fiber/v2"
)

func init() {

	isConfig := config.ConfigInit()
	isDatabase := database.InitDatabase()
	isRedisCache := database.InitRedisCache()
	if isConfig && isDatabase && isRedisCache {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database & Redis :: SUCCESS **** ", "Results", *database.Conn, database.RedisCaching.RedisClient)
	} else {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database & Redis :: ERROR **** ", "Results", *database.Conn, database.RedisCaching.RedisClient, logg.GetCallerPathNameFileNameLineNumber())
	}

}

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
