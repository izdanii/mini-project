package domain

import (
	"gorm.io/gorm"
)

type RepositoryVehiclesDB struct {
	db *gorm.DB
}

func NewVehiclesRepostoryDB(client *gorm.DB) RepositoryVehiclesDB {

	return RepositoryVehiclesDB{client}
}

type VehiclesRepositoryDB interface {
	FindAll() ([]Vehicles, error)
	FindByID(string) (Vehicles, error)
}

func (s *RepositoryVehiclesDB) FindAll() ([]Vehicles, error) {

	var vech []Vehicles
	err := s.db.Find(&vech).Error
	if err != nil {
		return vech, err
	}

	return vech, nil
}

func (s RepositoryVehiclesDB) FindByID(id string) (Vehicles, error) {

	var v Vehicles

	err := s.db.First(&v, id)

	if err != nil {
		return v, err.Error
	}

	return v, nil

}
