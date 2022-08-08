package domain

import (
	"fmt"

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
	DeleteID(string) (Vehicles, error)
	CreateID(Vehicles) (Vehicles, error)
	UpdateID(Vehicles, int) (Vehicles, error)
}

func (s *RepositoryVehiclesDB) FindAll() ([]Vehicles, error) {

	var vech []Vehicles
	fmt.Println("vech", vech)
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

func (s RepositoryVehiclesDB) DeleteID(id string) (Vehicles, error) {

	var v Vehicles

	err := s.db.Delete(&v, id)

	if err != nil {
		return v, err.Error
	}

	return v, nil

}

func (s RepositoryVehiclesDB) CreateID(vc Vehicles) (Vehicles, error) {

	err := s.db.Create(&vc).Error

	if err != nil {
		return vc, err
	}

	return vc, nil

}

func (s RepositoryVehiclesDB) UpdateID(vc Vehicles, id int) (Vehicles, error) {

	err := s.db.Model(&vc).Where("vehicle_id = ?", id).Updates(vc).Error

	if err != nil {
		return vc, err
	}

	return vc, nil

}
