package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type PlantController struct {
	plantService app.PlantService
}

func NewPlantController(ps app.PlantService) PlantController {
	return PlantController{
		plantService: ps,
	}
}

func (c PlantController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plant, err := requests.Bind(r, requests.PlantRequest{}, domain.Plant{})
		if err != nil {
			log.Printf("PlantController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		plant.UserId = user.Id
		plant, err = c.plantService.Save(plant)
		if err != nil {
			log.Printf("PlantController.Save(c.plantService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var resp resources.PlantDto
		resp = resp.DomainToDto(plant)
		Created(w, resp)
	}
}

func (c PlantController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		plants, err := c.plantService.FindList(user.Id)
		if err != nil {
			log.Printf("PlantController.FindList(c.plantService.FindList): %s", err)
			InternalServerError(w, err)
			return
		}

		resp := resources.PlantDto{}.DomainToDtoCollection(plants)
		Success(w, resp)
	}
}
