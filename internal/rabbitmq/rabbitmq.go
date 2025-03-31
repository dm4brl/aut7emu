package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Connect устанавливает подключение к RabbitMQ
func Connect(rabbitMQHost string) *amqp.Channel {
	conn, err := amqp.Dial(rabbitMQHost)
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка открытия канала RabbitMQ: %v", err)
	}

	return channel
}

// Publish публикует сообщение в RabbitMQ
func Publish(channel *amqp.Channel, queueName string, message string) {
	_, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка объявления очереди: %v", err)
	}

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Ошибка публикации сообщения: %v", err)
	}
}