package main

import (
	"fmt"

	"github.com/harrizsb/simple-crud/database"
	"github.com/harrizsb/simple-crud/user"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(3000)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/user/:username", user.GetUser)
	app.Post("/user/", user.CreateUser)
	app.Put("/user/", user.UpdateUser)
	app.Delete("/user/:username", user.DeleteUser)

	app.Post("/user/login", user.Login)
}
