package server

import (
	"fmt"
	"gocms/services/admins/users/transport/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server short",
	Long:  `server run`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Server run")

		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World 👋!")
		})

		app.Get("/hello", func(c *fiber.Ctx) error {
			return c.SendString("Hi 👋!")
		})
		app.Get("/about", func(c *fiber.Ctx) error {
			return c.SendString("About page 👋!")
		})

		app.Get("/about", func(c *fiber.Ctx) error {
			return c.SendString("About page 👋!")
		})

		hdl := handlers.NewUserHandler()
		app.Get("/users", hdl.ListUserHdl())

		app.Get("/users/create", hdl.CreateUserHdl())

		// app.Get("/users/change", handlers.ChangeUserHdl())
		// app.Get("/users/delete", handlers.DeleteUserHdl())

		app.Listen(":3000")
	},
}
