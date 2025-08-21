package usecase

import (
	"service-b/internal/domain/model"
	"service-b/internal/repository"
	"time"
)

// SensorUsecase defines available operations
type SensorUsecase interface {
	Save(data *model.SensorData) error
	Query(id1 string, id2 int, start, end *time.Time, limit, offset int) ([]model.SensorData, error)
	UpdateByID(id int64, newValue float64) error
	DeleteByID(id int64) error
}

type sensorUsecase struct {
	repo repository.SensorRepository
}

// NewSensorUsecase creates a new instance of usecase
func NewSensorUsecase(repo repository.SensorRepository) SensorUsecase {
	return &sensorUsecase{repo: repo}
}

func (u *sensorUsecase) Save(data *model.SensorData) error {
	return u.repo.Save(data)
}

func (u *sensorUsecase) Query(id1 string, id2 int, start, end *time.Time, limit, offset int) ([]model.SensorData, error) {
	// apply default limit
	if limit <= 0 {
		limit = 10
	}
	// apply default offset
	if offset < 0 {
		offset = 0
	}
	return u.repo.Query(id1, id2, start, end, limit, offset)
}

func (u *sensorUsecase) UpdateByID(id int64, newValue float64) error {
	return u.repo.UpdateByID(id, newValue)
}

func (u *sensorUsecase) DeleteByID(id int64) error {
	return u.repo.DeleteByID(id)
}
