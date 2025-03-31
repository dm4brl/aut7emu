package kafka

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

// Подключение к Kafka
func Connect(kafkaAddr string, topic string) *kafka.Writer {
    writer := &kafka.Writer{
        Addr:     kafka.TCP(kafkaAddr),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }

    // Проверка подключения
    _, err := writer.WriteMessages(context.Background(), kafka.Message{
        Value: []byte("test"),
    })
    if err != nil {
        log.Fatalf("Ошибка подключения к Kafka: %v", err)
    }

    return writer
}

// Отправка сообщения в Kafka
func SendMessage(writer *kafka.Writer, message string) {
    err := writer.WriteMessages(nil, kafka.Message{
        Value: []byte(message),
    })
    if err != nil {
        log.Printf("Ошибка отправки сообщения в Kafka: %v", err)
    }
}

