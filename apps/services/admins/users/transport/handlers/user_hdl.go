package handlers

import "github.com/gofiber/fiber/v2"

type userHdl struct {
}

func NewUserHandler() *userHdl {
	return &userHdl{}
}

func (h *userHdl) ListUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("list Users")
	}
}

func (h *userHdl) CreateUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Create Users")
	}
}

// func ChangeUserHdl() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return c.SendString("Change Users")
// 	}
// }
// func DeleteUserHdl() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		return c.SendString("Delete Users")
// 	}
// }
