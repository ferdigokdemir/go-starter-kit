package main

import (
	"go-starter-kit/config/database"
	"go-starter-kit/config/routes"

	"log"

	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
	"github.com/gofiber/recover"

	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/logger"

	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
)

// init is loading environment files.
func init() {
	viper.SetConfigFile(`./config/config.json`)
	error := viper.ReadInConfig()
	if error != nil {
		panic(error)
	}
}

func main() {
	app := fiber.New()

	if viper.GetBool("redisCache.enabled") {
		redis.Connect()
		app.Use(middlewares.RedisCache)
	}

	app.Use(recover.New())

	if viper.GetBool("compression.enabled") {
		app.Use(compression.New())
	}

	if viper.GetBool("cors.enabled") {
		app.Use(cors.New())
	}

	if viper.GetBool("logger.enabled") {
		app.Use(logger.New())
	}

	if viper.GetBool("prefork.enabled") {
		app.Settings.Prefork = true
	}

	if viper.GetBool("helmet.enabled") {
		app.Use(helmet.New())
	}

	if viper.GetBool("database.enabled") {
		database.Connect()
	}

	if viper.GetBool("favicon.ignore") {
		app.Use(middleware.Favicon())
	}

	if viper.GetBool("favicon.cache") {
		app.Use(middleware.Favicon("./public/favicon.ico"))
	}

	if viper.GetBool("ratelimit.enabled") {
		app.Use(limiter.New())
	}

	app.Static("/", "./public")

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(viper.GetInt("app.port")))
}
