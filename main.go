package main

import (
	"go-starter-kit/config/database"
	"go-starter-kit/config/routes"

	"github.com/gofiber/helmet"
	"log"

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

	if viper.GetBool("compression.enabled") {
		app.Use(compression.New(compression.Config{Level: compression.LevelBrotliBestCompression}))
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
		database.MongoConnect()
	}

	app.Static("/", "./public")

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(viper.GetInt("app.port")))
}
