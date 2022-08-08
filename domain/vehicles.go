package domain

type Vehicles struct {
	ID    int    `json:"vehicle_id" gorm:"column:vehicle_id"`
	Name  string `json:"name" gorm:"column:name"`
	Type  string `json:"type" gorm:"column:type"`
	Plat  string `json:"plat" gorm:"column:plat"`
	Color string `json:"color" gorm:"column:color"`
}

type InputVehicles struct {
	Name  string `json:"name" gorm:"column:name"`
	Type  string `json:"type" gorm:"column:type"`
	Plat  string `json:"plat" gorm:"column:plat"`
	Color string `json:"color" gorm:"column:color"`
}

type UpdateVehicles struct {
	ID    int    `json:"vehicle_id" gorm:"column:vehicle_id"`
	Name  string `json:"name" gorm:"column:name"`
	Type  string `json:"type" gorm:"column:type"`
	Plat  string `json:"plat" gorm:"column:plat"`
	Color string `json:"color" gorm:"column:color"`
}

type VehiclesRepository interface {
	FindAll() ([]Vehicles, error)
	FindByID() (Vehicles, error)
	DeleteID(string) (Vehicles, error)
	CreateID(Vehicles) (Vehicles, error)
	UpdateID(Vehicles) (Vehicles, error)
}
