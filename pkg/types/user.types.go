package types

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Pets Pets   `json:"pets"`
}

type Users []User

type UserCreateDto struct {
	Name string `json:"name"`
}

type UserResponseSerializer struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Pets string `json:"pets"`
}

type UsersResponseSerializer []UserResponseSerializer

// model

// entity

// dto
// serializer
