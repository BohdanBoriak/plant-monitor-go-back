package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const PlantsTableName = "plants"

type plant struct {
	Id        uint64           `db:"id,omitempty"`
	UserId    uint64           `db:"user_id"`
	Name      string           `db:"name"`
	City      string           `db:"city"`
	Address   string           `db:"address"`
	Type      domain.PlantType `db:"type"`
	CreatedAt time.Time        `db:"created_at"`
	UpdatedAt time.Time        `db:"updated_at"`
	DeletedAt *time.Time       `db:"deleted_at"`
}

type PlantRepository interface {
	Save(p domain.Plant) (domain.Plant, error)
}

type plantRepository struct {
	coll db.Collection
	sess db.Session
}

func NewPlantRepository(sess db.Session) PlantRepository {
	return plantRepository{
		coll: sess.Collection(PlantsTableName),
		sess: sess,
	}
}

func (r plantRepository) Save(p domain.Plant) (domain.Plant, error) {
	pl := r.mapDomainToModel(p)
	pl.CreatedAt = time.Now()
	pl.UpdatedAt = time.Now()

	err := r.coll.InsertReturning(&pl)
	if err != nil {
		return domain.Plant{}, err
	}

	p = r.mapModelToDomain(pl)
	return p, nil
}

func (r plantRepository) mapDomainToModel(p domain.Plant) plant {
	return plant{
		Id:        p.Id,
		UserId:    p.UserId,
		Name:      p.Name,
		City:      p.City,
		Address:   p.Address,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func (r plantRepository) mapModelToDomain(p plant) domain.Plant {
	return domain.Plant{
		Id:        p.Id,
		UserId:    p.UserId,
		Name:      p.Name,
		City:      p.City,
		Address:   p.Address,
		Type:      p.Type,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
