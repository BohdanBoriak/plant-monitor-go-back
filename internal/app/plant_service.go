package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type PlantService interface {
	Save(p domain.Plant) (domain.Plant, error)
}

type plantService struct {
	plantRepo database.PlantRepository
}

func NewPlantService(pr database.PlantRepository) PlantService {
	return plantService{
		plantRepo: pr,
	}
}

func (s plantService) Save(p domain.Plant) (domain.Plant, error) {
	plant, err := s.plantRepo.Save(p)
	if err != nil {
		log.Printf("plantService.Save(s.plantRepo.Save): %s", err)
		return domain.Plant{}, err
	}

	return plant, nil
}
