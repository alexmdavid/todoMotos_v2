package models

import "time"

type Blog struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Titulo          string    `gorm:"type:text;not null"`
	Contenido       string    `gorm:"type:text;not null"`
	FechaPublicacion time.Time `gorm:"not null"`
	Autor           string    `gorm:"type:text;not null"`
	Resumen         string    `gorm:"type:text"`
	ImagenUrl       string    `gorm:"type:text"`
	Etiquetas       string    `gorm:"type:text"`
}
