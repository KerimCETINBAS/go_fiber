package rest

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/domain/pet"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/domain/user"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/rest/handlers"
)

func New() *fiber.App {

	userService := user.UserService()
	petService := pet.PetService()
	app := fiber.New()

	app.Use(swagger.New(swagger.Config{
		FilePath: "docs/swagger.json",
	}))

	handlers.UserHandler(app.Group("/users"), &userService)
	handlers.PetHandler(app.Group("/pets"), &petService)
	return app
}
