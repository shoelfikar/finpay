package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/shoelfikar/API/model"
)


type Router struct {
	Router *fiber.App
}

type version struct {
	Version string `json:"version"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func (r *Router) Setup(baseURL string) {
	// init router config
	r.Router = fiber.New(fiber.Config{
		AppName: "Finpay - Realtime Transaction",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseJSON{
				Error: "Internal Server Error",
				Success: "false",
				Data: struct{}{},
			})
		},
	})

	r.Router.Use(recover.New())

	// logging config
	r.Router.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// cors config
	r.Router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token",
		MaxAge: 86400,
		// AllowCredentials: false,
		ExposeHeaders: "Content-Length",
	}))

	r.Router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(version{
			Name: "Finpay - Realtime Transaction",
			Description: "Finpay E Wallet REST API",
			Version: "v1.0",
		})
	})

	r.setupAPIRouter(baseURL)
}

func (r *Router) setupAPIRouter(baseURL string) {
	apiGroup := r.Router.Group(baseURL)

	r.configureAuthRouter(apiGroup)
}

func (r *Router) configureAuthRouter(router fiber.Router) {
	authGroup := router.Group("/auth")
	authGroup.Post("/login", func (c *fiber.Ctx) error {
		return c.SendString("success request login")
	})
}

func (r *Router) Run(port string) {
	if r.Router == nil {
		panic("[ROUTER ERROR] Server has not been initialized. Make sure to call Setup() before Run().")
	}

	err := r.Router.Listen(":" + port)
	if err != nil {
		panic(fmt.Sprintf("[SERVER ERROR] Failed to start the server on port %s: %v", port, err))
	}
}