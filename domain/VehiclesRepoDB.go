package domain

import (
	"mini-project/dto"

	"gorm.io/gorm"
)

type RepositoryVehiclesDB struct {
	db *gorm.DB
}

func NewVehiclesRepostoryDB(client *gorm.DB) RepositoryVehiclesDB {

	return RepositoryVehiclesDB{client}
}

type VehiclesRepositoryDB interface {
	FindAll(dto.Pagination) (dto.Pagination, error)
	FindByID(string) (Vehicles, error)
	DeleteID(string) (Vehicles, error)
	CreateID(Vehicles) (Vehicles, error)
	UpdateID(Vehicles, int) (Vehicles, error)
}

func (s *RepositoryVehiclesDB) FindAll(pagination dto.Pagination) (dto.Pagination, error) {
	var p dto.Pagination
	tr := 0
	offset := pagination.Page * pagination.Limit
	var vech []Vehicles
	errFind := s.db.Limit(pagination.Limit).Offset(offset).Find(&vech).Error
	if errFind != nil {
		return p, nil
	}
	pagination.Rows = vech
	totalRows := int64(tr)
	errCount := s.db.Model(vech).Count(&totalRows).Error
	if errCount != nil {
		return p, errCount
	}

	return pagination, nil
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
