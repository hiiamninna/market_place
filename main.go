package main

import (
	"encoding/json"
	"fmt"
	"market_place/collections"
	"market_place/controller"
	"market_place/library"
	"market_place/repository"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Use recover middleware to prevent crashes
	app.Use(recover.New())
	// Register the Prometheus metrics handler
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	// use this route to test
	/**
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "running an api at port 8000",
		})
	})
	app.Get("/test", context.JWT.Authentication(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "you are authorized",
		})
	})
	**/

	user := app.Group("/v1/user")
	{
		user.Post("/register", ParseContext(context.CTL.USER.Register))
		user.Post("/login", ParseContext(context.CTL.USER.Login))
	}

	product := app.Group("/v1/product")
	{
		product.Post("", context.JWT.Authentication(), ParseContext(context.CTL.PRODUCT.Create))
		product.Patch("/:id", context.JWT.Authentication(), ParseContext(context.CTL.PRODUCT.Update))
		product.Delete("/:id", context.JWT.Authentication(), ParseContext(context.CTL.PRODUCT.Delete))

		product.Post("/:id/buy", context.JWT.Authentication(), ParseContext(context.CTL.PAYMENT.Create))
		product.Post("/:id/stock", context.JWT.Authentication(), ParseContext(context.CTL.PRODUCT.UpdateStock))

		product.Get("", ParseContextList(context.CTL.PRODUCT.List))
		product.Get("/:id", ParseContext(context.CTL.PRODUCT.Get))

		// Simplified route creation with NewRoute function, added to prometheus-grafana, monitoring
		NewRoute(app, "/pg/v1/product/", "GET", ParseContextList(context.CTL.PRODUCT.List))
		NewRoute(app, "/pg/v1/product/:id", "GET", ParseContext(context.CTL.PRODUCT.Get))
	}

	image := app.Group("/v1/image")
	{
		image.Post("", context.JWT.Authentication(), ParseContext(context.CTL.IMAGE.ImageUpload))
	}

	bankAccount := app.Group("/v1/bank/account")
	{
		bankAccount.Post("", context.JWT.Authentication(), ParseContext(context.CTL.BANK_ACCOUNT.Create))
		bankAccount.Patch("/:id", context.JWT.Authentication(), ParseContext(context.CTL.BANK_ACCOUNT.Update))
		bankAccount.Delete("/:id", context.JWT.Authentication(), ParseContext(context.CTL.BANK_ACCOUNT.Delete))
		bankAccount.Get("", context.JWT.Authentication(), ParseContext(context.CTL.BANK_ACCOUNT.List))
	}

	return app
}

func ParseContext(f func(*fiber.Ctx) (int, string, interface{}, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		code, message, resp, err := f(c)
		c.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:01:02 "), err)
			errBody := Response(message, nil)
			c.Set("Content-Length", fmt.Sprintf("%d", len(errBody)))
			return c.Status(code).Send(errBody)
		}

		successBody := Response(message, resp)
		c.Set("Content-Length", fmt.Sprintf("%d", len(successBody)))
		return c.Status(code).Send(successBody)
	}
}

func ParseContextList(f func(*fiber.Ctx) (int, string, collections.Meta, interface{}, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		code, message, meta, resp, err := f(c)
		c.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:01:02 "), err)
			errBody := Response(message, nil)
			c.Set("Content-Length", fmt.Sprintf("%d", len(errBody)))
			return c.Status(code).Send(errBody)
		}

		successBody := ResponseList(message, meta, resp)
		c.Set("Content-Length", fmt.Sprintf("%d", len(successBody)))
		return c.Status(code).Send(successBody)
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

func ResponseList(message string, meta collections.Meta, data interface{}) []byte {

	response := struct {
		Message string           `json:"message"`
		Data    interface{}      `json:"data"`
		Meta    collections.Meta `json:"meta"`
	}{
		Message: message,
		Data:    data,
		Meta:    meta,
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
	JWT library.JWT
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

	s3, err := library.NewS3(config.S3Config)
	if err != nil {
		fmt.Println("new s3 : %w", err)
	}

	// set up repo and controller
	repo := repository.NewRepository(db)
	ctl := controller.NewController(repo, jwt, config.BcryptSalt, s3)

	return Context{
		CFG: config,
		JWT: jwt,
		CTL: ctl,
		S3:  s3,
	}, nil
}

// PROMETHEUS n GRAFANA
var (
	helloRequestHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "hiiamninna_request",
		Help:    "Histogram of the /hiiamninna request duration.",
		Buckets: prometheus.LinearBuckets(1, 1, 10), // Adjust bucket sizes as needed
	}, []string{"path", "method", "status"})
)

func NewRoute(app *fiber.App, path string, method string, handler fiber.Handler) {
	app.Add(method, path, wrapHandlerWithMetrics(path, method, handler))
}

func wrapHandlerWithMetrics(path string, method string, handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		// Execute the actual handler and catch any errors
		err := handler(c)

		// Regardless of whether an error occurred, record the metrics
		duration := time.Since(startTime).Seconds()
		statusCode := fmt.Sprintf("%d", c.Response().StatusCode())
		if err != nil {
			if c.Response().StatusCode() == fiber.StatusOK { // Default status code
				statusCode = "500" // Assume internal server error if not set
			}
			c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // Ensure the response reflects the error
		}

		helloRequestHistogram.WithLabelValues(path, method, statusCode).Observe(duration)
		return err
	}
}
