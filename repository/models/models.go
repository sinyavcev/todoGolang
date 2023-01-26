package models

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"autoIncrement" gorm:"primaryKey: type:uuid"`
	Title     string    `json:"title" gorm:"unique"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
