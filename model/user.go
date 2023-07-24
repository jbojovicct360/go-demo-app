package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint `gorm:"primaryKey"`
	Username string
	Tasks    []Task
}
