package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	user "github.com/kerimcetinbas/fiber_paseto_casbin/pkg/domain/user"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/types"
)

type userHandler struct {
	userService user.IUserService
}

func UserHandler(
	router fiber.Router,
	userService *user.IUserService) {
	uh := &userHandler{
		userService: *userService,
	}
	router.Get("/", uh.GetUsers)
	router.Post("/", uh.CreateUser)
	router.Delete("/:id", uh.DelteUserById)
}

// Get all users godoc
//
//	@Summary		get all userss
//	@Description	get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/accounts/{id} [get]
func (h *userHandler) GetUsers(ctx *fiber.Ctx) error {
	fmt.Println("Get users")
	users, _ := h.userService.GetUsers()

	response := make(types.UsersResponseSerializer, 0)
	for _, u := range users {
		response = append(response, types.UserResponseSerializer{
			Id:   u.Id,
			Name: u.Name,
			Pets: fmt.Sprintf("http://localhost:8082/pets/user/%v", u.Id),
		})
	}
	ctx.JSON(&response)
	return nil
}

func (h *userHandler) CreateUser(ctx *fiber.Ctx) error {
	var userData types.User
	ctx.BodyParser(&userData)
	h.userService.CreateUser(userData)
	return nil
}

func (h *userHandler) DelteUserById(ctx *fiber.Ctx) error {

	idx, _ := strconv.ParseUint(ctx.Params("id"), 10, 32)

	h.userService.DeleteUserById(uint(idx))
	return nil
}
