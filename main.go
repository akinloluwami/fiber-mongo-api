package main

import (
	"fiber-mongo-api/configs"

	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
	return c.JSON((&fiber.Map{"data":"Hello from Fiber and MongoDB"}))
})
configs.ConnectDB()
routes.UserRoute(app)
app.Listen(":1600")
}