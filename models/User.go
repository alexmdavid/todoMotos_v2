package models


type User struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Nombre          string `gorm:"type:varchar(100);not null"`
	Correo          string `gorm:"type:varchar(160);not null"`
	ClaveHash       string `gorm:"type:text;not null"`
	Rol             string `gorm:"type:text;default:''"`
	ImagenPerfilUrl string `gorm:"type:text"`
}

func (User) TableName() string {
	return "Users"
}