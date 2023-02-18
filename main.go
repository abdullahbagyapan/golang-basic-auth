package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"net/http"
)

func main() {

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{ // filter process
		Users: map[string]string{ // registered accounts
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool { // checking
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusUnauthorized) // if checking return false do ...
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/home", func(ctx *fiber.Ctx) error { // after filter process
		return ctx.JSON("hello world")
	})

	app.Listen(":8080")

}
