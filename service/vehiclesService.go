package service

import (
	"fmt"
	"mini-project/domain"
)

type VehiclesService interface {
	GetAllVehicles() ([]domain.Vehicles, error)
	GetVehiclesByID(string) (*domain.Vehicles, error)
	DeleteVehiclesByID(string) (*domain.Vehicles, error)
	CreateVehiclesByID(domain.InputVehicles) (domain.Vehicles, error)
	UpdateVehiclesByID(domain.UpdateVehicles, int) (domain.Vehicles, error)
}

type DefaultVehiclesService struct {
	repo domain.VehiclesRepositoryDB
}

func (s DefaultVehiclesService) GetAllVehicles() ([]domain.Vehicles, error) {
	vehicles, err := s.repo.FindAll()
	if err != nil {
		return vehicles, err
	}

	return vehicles, nil
}

func (s DefaultVehiclesService) GetVehiclesByID(id string) (*domain.Vehicles, error) {
	v, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// response := v.domain.Vehicles
	return &v, nil
}

func (s DefaultVehiclesService) DeleteVehiclesByID(id string) (*domain.Vehicles, error) {
	v, err := s.repo.DeleteID(id)
	if err != nil {
		return nil, err
	}

	// response := v.domain.Vehicles
	return &v, nil
}

func (s DefaultVehiclesService) CreateVehiclesByID(input domain.InputVehicles) (domain.Vehicles, error) {
	vc := domain.Vehicles{}
	fmt.Println(input)
	vc.Name = input.Name
	vc.Type = input.Type
	vc.Plat = input.Plat
	vc.Color = input.Color

	v, err := s.repo.CreateID(vc)
	if err != nil {
		return v, err
	}

	// response := v.domain.Vehicles
	return v, nil
}

func (s DefaultVehiclesService) UpdateVehiclesByID(update domain.UpdateVehicles, id int) (domain.Vehicles, error) {

	vc := domain.Vehicles{}
	fmt.Println(update)
	vc.ID = id
	vc.Name = update.Name
	vc.Type = update.Type
	vc.Plat = update.Plat
	vc.Color = update.Color

	v, err := s.repo.UpdateID(vc, id)
	if err != nil {
		return v, err
	}

	// response := v.domain.Vehicles
	return v, nil
}

func NewVehiclesService(repository domain.VehiclesRepositoryDB) DefaultVehiclesService {
	return DefaultVehiclesService{repository}
}
