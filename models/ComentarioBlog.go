package models

import "time"

type ComentarioBlog struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Contenido       string    `gorm:"type:text;not null"`
	Likes           int       `gorm:"not null"`
	UnLikes         int       `gorm:"not null"`
	FechaComentario time.Time `gorm:"not null"`
	BlogID          uint      `gorm:"not null"`
	UsuarioID       uint      `gorm:"not null"`
}
