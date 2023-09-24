package user

import (
	"errors"
	"fmt"

	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/storage"
	"github.com/kerimcetinbas/fiber_paseto_casbin/pkg/types"
)

type IUserService interface {
	GetUsers() (types.Users, error)
	CreateUser(types.User) error
	DeleteUserById(id uint) error
}
type userService struct{}

func UserService() IUserService { return &userService{} }

func (us *userService) GetUsers() (types.Users, error) {

	users := make(types.Users, 0)
	if u, ok := storage.Memory["Users"].(types.Users); ok {

		for i, user := range u {
			user.Pets = make(types.Pets, 0)

			if p, o := storage.Memory["Pets"].(types.Pets); o {

				for _, pet := range p {
					if pet.Owner == user.Id {
						fmt.Println(pet)

						user.Pets = append(user.Pets, pet)

						fmt.Println(u)
					}
				}
			}

			u[i] = user

		}

		users = u
	}
	return users, nil
}

func (us *userService) CreateUser(user types.User) error {
	fmt.Println(user)
	if u, ok := storage.Memory["Users"].(types.Users); ok {

		user.Id = uint(len(u))
		storage.Memory["Users"] = append(u, user)
	} else {
		fmt.Println(user)
		users := make(types.Users, 0)
		users = append(users, user)
		storage.Memory["Users"] = users
	}
	return nil
}

func (us *userService) DeleteUserById(id uint) error {

	if u, ok := storage.Memory["Users"].(types.Users); ok {

		for index, user := range u {

			if user.Id == id {
				storage.Memory["Users"] = append(u[:index], u[index+1:]...)
			}
		}
	} else {
		return errors.New("Not found")
	}
	return nil
}
