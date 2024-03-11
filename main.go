package main

import (
	"encoding/json"
	"fmt"
	"market_place/controller"
	"market_place/library"
	"market_place/repository"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := route()

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

func route() *fiber.App {

	context, _ := NewContext()

	// set route
	app := fiber.New(fiber.Config{})

	// use this route to test
	/**
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "running an api at port 8000",
		})
	})
	app.Get("/test", jwt.Authentication(config.JWTSecret), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "you are authorized",
		})
	})
	**/
	app.Post("/v1/user/register", ParseContext(context.CTL.USER.Register))

	return app
}

func ParseContext(f func(*fiber.Ctx) (int, string, interface{}, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		code, message, resp, err := f(c)
		c.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
			errBody := Response(message, nil)
			c.Set("Content-Length", fmt.Sprintf("%d", len(errBody)))
			return c.Status(code).Send(errBody)
		}

		successBody := Response(message, resp)
		c.Set("Content-Length", fmt.Sprintf("%d", len(successBody)))
		return c.Send(Response(message, resp))
	}
}

func Response(message string, data interface{}) []byte {

	response := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: message,
		Data:    data,
	}

	value, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	}

	return value
}

type Context struct {
	CFG library.Config
	CTL controller.Controller
	S3  library.S3
}

func NewContext() (Context, error) {
	// read config
	config, err := library.NewConfiguration()
	if err != nil {
		fmt.Println("new config : %w", err)
	}

	// set up jwt
	jwt := library.NewJWT(config.JWTSecret)

	// set up db
	db, err := library.NewDatabaseConnection(config.DB)
	if err != nil {
		fmt.Println("new db : %w", err)
	}

	// set up repo and controller
	repo := repository.NewRepository(db)
	ctl := controller.NewController(repo, jwt, config.BcryptSalt)

	return Context{
		CFG: config,
		CTL: ctl,
		S3:  library.S3{},
	}, nil
}
