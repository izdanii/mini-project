package domain

type Vehicles struct {
	ID    string `json:"vehicle_id" gorm:"column:vehicle_id"`
	Name  string `json:"name" gorm:"column:name"`
	Type  string `json:"type" gorm:"column:type"`
	Plat  string `json:"plat" gorm:"column:plat"`
	Color string `json:"color" gorm:"column:color"`
}

type VehiclesRepository interface {
	FindAll() ([]Vehicles, error)
	FindByID() (Vehicles, error)
}
