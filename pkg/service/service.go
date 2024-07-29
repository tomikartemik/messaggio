package service

import (
	"github.com/segmentio/kafka-go"
	"messaggio/pkg/model"
	"messaggio/pkg/repository"
)

type Service struct {
	Message
}

func NewService(repos *repository.Repository, kafkaWriter *kafka.Writer) *Service {
	return &Service{
		Message: NewMessageService(repos.Message, kafkaWriter),
	}
}

type Message interface {
	SaveNewMessage(message model.Message) error
	ProcessedMessagesStats() (int, error)
}
