package models

import "time"

type Review struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Calificacion int     `gorm:"not null"`
	Comentario string    `gorm:"type:text;not null"`
	Fecha      time.Time `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	BikeID     uint      `gorm:"not null"`
}
