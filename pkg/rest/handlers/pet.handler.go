package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/domain/pet"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/types"
)

type petHandler struct {
	petService pet.IPetService
}

func PetHandler(
	router fiber.Router,
	petService *pet.IPetService) {

	ph := petHandler{
		petService: *petService,
	}

	router.Get("/", ph.GetPets)
	router.Post("/", ph.CreaePet)
	router.Get("/user/:id", ph.GetPetByUserId)
}

func (ph *petHandler) GetPets(ctx *fiber.Ctx) error {

	pets, _ := ph.petService.GetPets()

	ctx.JSON(&pets)
	return nil
}

func (ph *petHandler) CreaePet(ctx *fiber.Ctx) error {

	pet := new(types.PetCreateDto)
	ctx.BodyParser(&pet)

	ph.petService.CreatePet(*pet)
	return nil
}

func (ph *petHandler) GetPetByUserId(ctx *fiber.Ctx) error {
	idx, _ := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if pets, err := ph.petService.GetPetByUserId(uint(idx)); err != nil {

		return fiber.ErrNotFound
	} else {
		ctx.JSON(&pets)
	}

	return nil
}
