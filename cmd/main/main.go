package main

import (
	"log"
	"os"

	"github.com/DewaBiara/Secure-DOCS/pkg/bootsrapper"
	"github.com/DewaBiara/Secure-DOCS/pkg/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/DewaBiara/Secure-DOCS/pkg/database"
)

func init() {
	if os.Getenv("ENV") == "production" {
		return
	}

	//	load env variables from .env file for local development
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	env := config.LoadConfig()

	db, err := database.Connect(
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_USER"],
		env["DB_PASS"],
		env["DB_NAME"],
		5,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatalf(err.Error())
	}

	e := echo.New()
	bootsrapper.InitController(e, db, env)

	e.Logger.Fatal(e.Start(":" + env["PORT"]))
}
