package models

import (
	"time"
)

type Message struct {
	ID             uint64       `gorm:"column:id" json:"id"`
	Message        string      	`gorm:"column:message" json:"message"`
	CreatedAt      time.Time   	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time   	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      *time.Time  	`gorm:"column:deleted_at" json:"deleted_at"`
}

func (Message) TableName() string {
	return "messages"
}
