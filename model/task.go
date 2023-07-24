package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	id          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint `gorm:"index"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
