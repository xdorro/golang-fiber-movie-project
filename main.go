package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/pkg/config"
	"github.com/xdorro/golang-fiber-movie-project/pkg/middleware"
	"github.com/xdorro/golang-fiber-movie-project/pkg/router"
	"log"
	"os"
	"os/signal"
	"syscall"

	// docs are generated by Swag CLI, you have to import them.
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
)

// @title Golang Fiber Base Project
// @version 1.0.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Tuan Anh Nguyen Manh
// @contact.email xdorro@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	serverConfig := config.GetServer()

	app := fiber.New(fiber.Config{
		AppName: serverConfig.Name,

		Prefork: serverConfig.Prefork,

		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 statusCode
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}
			// Set Content-Type: application/json; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			// Return statusCode with error message
			return c.Status(code).Send([]byte(err.Error()))
		},
	})

	// Attach Middlewares.
	middleware.BaseMiddleware(app)

	// Routes.
	router.BaseRouter(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Printf("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	if err := app.Listen(serverAddr); err != nil {
		log.Printf("Oops... server is not running! error: %v", err)
	}
}
