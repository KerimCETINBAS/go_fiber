package pet

import (
	"errors"

	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/storage"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/types"
)

type IPetService interface {
	GetPets() (types.Pets, error)
	CreatePet(types.PetCreateDto) error
	DeletePetById(id uint) error
	GetPetByUserId(id uint) (types.Pets, error)
}
type petService struct{}

func PetService() IPetService { return &petService{} }

func (ps *petService) CreatePet(pet types.PetCreateDto) error {
	p := types.Pet{
		Name:  pet.Name,
		Owner: pet.Owner,
	}
	if u, ok := storage.Memory["Pets"].(types.Pets); ok {

		storage.Memory["Pets"] = append(u, p)
	} else {
		pets := make(types.Pets, 0)
		pets = append(pets, p)
		storage.Memory["Pets"] = pets
	}
	return nil
}

func (ps *petService) GetPets() (types.Pets, error) {

	if p, ok := storage.Memory["Pets"].(types.Pets); ok {
		return p, nil
	}

	return make(types.Pets, 0), nil

}

func (ps *petService) DeletePetById(id uint) error {

	return nil
}

func (ps *petService) GetPetByUserId(id uint) (types.Pets, error) {
	pe := make(types.Pets, 0)
	isExist := false
	if u, ok := storage.Memory["Users"].(types.Users); ok {

		for _, us := range u {

			if us.Id != id {
				continue
			} else {
				isExist = true
			}
		}

	}

	if !isExist {
		return pe, errors.New("User does not exist")
	}
	if p, ok := storage.Memory["Pets"].(types.Pets); ok {
		for _, pet := range p {
			if pet.Owner == id {
				pe = append(pe, pet)
			}
		}

		return pe, nil

	}
	return make(types.Pets, 0), nil
}
