package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"taheri24.ir/publish-server/database"
	"taheri24.ir/publish-server/internal"
	"taheri24.ir/publish-server/routers"
)

func main() {
	internal.ShowBanner("Publish Server v1.0")
	db, err := database.InitDbConnection()
	if err != nil {
		panic("failed to connect database")
	}
	database.Connection = db
	config := fiber.Config{
		DisableStartupMessage: true,
	}

	app := fiber.New(config)
	database.SyncDatabaseTables(db)
	routers.SetupApiRouters(app)
	listenTo := "0.0.0.0:3300"
	fmt.Println("Server Started at  " + listenTo)
	if err := app.Listen(listenTo); err != nil {
		panic(err)
	}

}
