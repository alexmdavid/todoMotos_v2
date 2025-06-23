package models

import "time"

type ComentarioBlogReaction struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	ComentarioID uint      `gorm:"not null"`
	UserID       uint      `gorm:"not null"`
	Tipo         int       `gorm:"not null"`
	Fecha        time.Time `gorm:"not null"`
}

func  (ComentarioBlogReaction) TableName() string {
	return "ComentariosBlogReaction"
}
