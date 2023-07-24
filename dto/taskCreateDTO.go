package dto

type CreateTaskDTO struct {
	Title       string
	Description string
	UserID      uint `gorm:"index"`
}
