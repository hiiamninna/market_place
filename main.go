package main

import (
	"fmt"
	"market_place/library"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_, err := library.NewConfiguration()
	if err != nil {
		fmt.Println("new config : %w", err)
	}

	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "running an api at port 8000",
		})
	})

	//set channel to notify when app interrupted
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":8000"); err != nil {
		fmt.Println("error on http listen : %w", err)
	}
}
