package service

import "mini-project/domain"

type VehiclesService interface {
	GetAllVehicles() ([]domain.Vehicles, error)
	GetVehiclesByID(string) (*domain.Vehicles, error)
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

func NewVehiclesService(repository domain.VehiclesRepositoryDB) DefaultVehiclesService {
	return DefaultVehiclesService{repository}
}
