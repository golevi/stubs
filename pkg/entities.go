package pkg

import (
	"github.com/golevi/mocktest/pkg/pet"
	"github.com/golevi/mocktest/pkg/user"
)

type Entities interface {
	GetUser(id string) (user.User, error)
	GetPets(userID string) ([]pet.Pet, error)
}
