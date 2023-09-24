package types

type Pet struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Owner uint   `json:"owner"`
}

type Pets []Pet

type PetCreateDto struct {
	Name  string `json:"name"`
	Owner uint   `json:"owner"`
}
