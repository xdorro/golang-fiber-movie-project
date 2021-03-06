package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

func BaseRouter(app *fiber.App) {
	app.Static("/storage", "./storage")

	// Create route group.
	app.Get("/swagger/*", swagger.Handler)

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":  "Welcome to Fiber Go API!",
			"docs": "/swagger/index.html",
		})
	})

	api.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Welcome to Fiber Go API!",
		})
	})

	// Private router
	privateRoute(api)

	// Public router
	publicRoute(api)

	// Auth Router
	authRouter(api)

	// 404 Not Found Router
	notFoundRoute(app)
}

func notFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return util.ResponseNotFound("Đường dẫn không tồn tại")
		},
	)
}
