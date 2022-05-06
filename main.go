package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"main.go/user"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my app")
}
func Router(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}
func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/api", welcome)
	Router(app)
	log.Fatal(app.Listen(":3000"))
}
