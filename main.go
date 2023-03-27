package main

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter-unit-test/api/route"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/cache"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/database"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
	"github.com/gofiber/fiber/v2"
)

func init() {

	isConfig := config.ConfigInit()
	isDatabase := database.InitDatabase()
	isCache := cache.InitCache()
	if isConfig && isDatabase && isCache {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database & Redis :: SUCCESS **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient)
	} else {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database & Redis :: ERROR **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient, logg.GetCallerPathNameFileNameLineNumber())
		panic(fmt.Errorf("initail configuration error"))
	}

}

func main() {
	defer database.ConnectionClose()

	app := fiber.New()
	router := route.NewRoute()
	router.InitRoute(app)

	app.Listen(config.Env.API_PORT)
}
