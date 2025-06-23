package models

import "time"

type ReviewReaction struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	ReviewID uint      `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	Tipo     int       `gorm:"not null"`
	Fecha    time.Time `gorm:"not null"`
}

func (ReviewReaction) TableName() string {
	return "ReviewsReactions"
}