package app

import "github.com/golevi/mocktest/pkg"

type Logic struct {
	Entities pkg.Entities
}

func (l Logic) GetPetNames(userID string) ([]string, error) {
	pets, err := l.Entities.GetPets(userID)
	if err != nil {
		return nil, err
	}

	out := make([]string, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}

	return out, nil
}
