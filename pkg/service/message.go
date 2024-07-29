package service

import (
	"context"
	"github.com/segmentio/kafka-go"
	"messaggio/pkg/model"
	"messaggio/pkg/repository"
)

type MessageService struct {
	repo        repository.Message
	kafkaWriter *kafka.Writer
}

func NewMessageService(repo repository.Message, kafkaWriter *kafka.Writer) *MessageService {
	return &MessageService{repo: repo, kafkaWriter: kafkaWriter}
}

//Синхронный вариант

func (s *MessageService) SaveNewMessage(message model.Message) error {
	message.Processed = false
	messageId, err := s.repo.SaveNewMessage(message)
	if err != nil {
		return err
	}

	err = s.kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message.Text),
		},
	)

	if err != nil {
		return err
	}

	err = s.repo.MarkMessageAsProcessed(messageId)
	if err != nil {
		return err
	}

	return nil
}

//Асинхронный вариант

//func (s *MessageService) SaveNewMessage(message model.Message) error {
//	message.Processed = false
//	err := s.repo.SaveNewMessage(message)
//	if err != nil {
//		return err
//	}
//
//	go s.processMessageAsync(message)
//
//	return nil
//}
//
//func (s *MessageService) processMessageAsync(message model.Message) {
//	_, _, err := s.producer.SendMessage(&sarama.ProducerMessage{
//		Topic: "messages",
//		Value: sarama.StringEncoder(message.Text),
//	})
//	if err != nil {
//		log.Println("Error sending message to Kafka:", err)
//		return
//	}
//
//	err = s.repo.MarkMessageAsProcessed(message.ID)
//	if err != nil {
//		log.Println("Error marking message as processed:", err)
//	}
//}

func (s *MessageService) ProcessedMessagesStats() (int, error) {
	return s.repo.ProcessedMessagesStats()
}
