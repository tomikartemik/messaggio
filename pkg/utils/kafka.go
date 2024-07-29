package utils

import (
	"github.com/segmentio/kafka-go"
	"os"
)

func InitKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    os.Getenv("KAFKA_TOPIC"),
		Balancer: &kafka.LeastBytes{},
	}
}

func InitKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_BROKER")},
		Topic:    os.Getenv("KAFKA_TOPIC"),
		GroupID:  os.Getenv("KAFKA_GROUP"),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func InitKafkaTopic() {
	controllerConn, err := kafka.Dial("tcp", os.Getenv("KAFKA_BROKER"))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{{Topic: "messages", NumPartitions: 1, ReplicationFactor: 1}}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
