package repository

import (
	"gorm.io/gorm"
	"messaggio/pkg/model"
)

type MessagePostgres struct {
	db *gorm.DB
}

func NewMessagePostgres(db *gorm.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) SaveNewMessage(message model.Message) (int, error) {
	err := r.db.Create(&message).Error
	return message.ID, err
}

func (r *MessagePostgres) MarkMessageAsProcessed(id int) error {
	return r.db.Model(model.Message{}).Where("id = ?", id).Update("processed", true).Error
}

func (r *MessagePostgres) ProcessedMessagesStats() (int, error) {
	var totalNumberOfProcessedMessages int64
	err := r.db.Model(&model.Message{}).Where("processed = true").Count(&totalNumberOfProcessedMessages).Error
	return int(totalNumberOfProcessedMessages), err
}
