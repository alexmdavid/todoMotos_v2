package models

type ListaDeFavoritos struct {
	BikeID  uint `gorm:"primaryKey"`
	UserID  uint `gorm:"primaryKey"`
}

func (ListaDeFavoritos) TableName() string {
	return "ListasDeFavoritos"
}