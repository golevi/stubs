package app

import (
	"github.com/golevi/mocktest/pkg"
	"github.com/golevi/mocktest/pkg/pet"
	"log"
	"testing"
)

type GetPetNamesStub struct {
	pkg.Entities
}

func (ps GetPetNamesStub) GetPets(userID string) ([]pet.Pet, error) {
	return []pet.Pet{{"Levi"}, {"Elz"}, {"Frank"}}, nil
}

func TestLogic_GetPetNames(t *testing.T) {
	l := Logic{GetPetNamesStub{}}
	names, err := l.GetPetNames("1")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(names)
	if len(names) == 0 {
		t.Error("invalid length of names")
	}
}
