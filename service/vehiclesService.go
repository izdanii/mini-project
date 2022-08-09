package service

import (
	"fmt"
	"mini-project/domain"
	"mini-project/dto"
	"mini-project/setupDB"
)

type VehiclesService interface {
	GetAllVehicles(dto.Pagination, int) (dto.Pagination, error)
	GetVehiclesByID(string) (*domain.Vehicles, error)
	DeleteVehiclesByID(string) (*domain.Vehicles, error)
	CreateVehiclesByID(domain.InputVehicles) (domain.Vehicles, error)
	UpdateVehiclesByID(domain.UpdateVehicles, int) (domain.Vehicles, error)
}

type DefaultVehiclesService struct {
	repo domain.VehiclesRepositoryDB
}

func (s DefaultVehiclesService) GetAllVehicles(pag dto.Pagination, id int) (dto.Pagination, error) {
	var p dto.Pagination
	db, _ := setupDB.ClientDB()
	userRepo := domain.NewUsersRepositoryDB(db)
	user, err := userRepo.FindByID(id)
	if err != nil {
		return p, err
	}
	if user.Username == "" {
		return p, err
	}
	vehicles, err := s.repo.FindAll(pag)
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
