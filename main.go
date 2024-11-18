package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Ninja struct {
	Name   string
	Weapon string
}

var ninja Ninja

func getNinja(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(
			err.Error(),
		)
		return err
	}
	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}
	return ctx.Status(fiber.StatusCreated).JSON(ninja)
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(logger.New())
	app.Use(requestid.New())

	ninjaGroupApp := app.Group("/ninja")
	ninjaGroupApp.Get("/", getNinja)
	ninjaGroupApp.Post("/", createNinja)
	app.Listen(":3000")
}
