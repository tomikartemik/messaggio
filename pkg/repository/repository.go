package repository

import (
	"gorm.io/gorm"
	"messaggio/pkg/model"
)

type Repository struct {
	Message
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Message: NewMessagePostgres(db),
	}
}

type Message interface {
	SaveNewMessage(message model.Message) (int, error)
	MarkMessageAsProcessed(id int) error
	ProcessedMessagesStats() (int, error)
}
