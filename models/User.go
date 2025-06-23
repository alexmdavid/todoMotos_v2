package models


type User struct {
	ID              uint   `gorm:"column:Id;primaryKey;autoIncrement" json:"id"`
	Nombre          string `gorm:"type:varchar(100);not null" json:"nombre"`
	Correo          string `gorm:"type:varchar(160);not null" json:"correo"`
	ClaveHash       string `gorm:"type:text;not null" json:"claveHash"`
	Rol             string `gorm:"type:text;default:''" json:"rol"`
	ImagenPerfilUrl string `gorm:"type:text" json:"imagenPerfilUrl"`
}


func (User) TableName() string {
	return "Users"
}