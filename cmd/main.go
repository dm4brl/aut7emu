package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dm4brl/aut7emu/internal/db"
	"github.com/dm4brl/aut7emu/internal/devices"
	"github.com/dm4brl/aut7emu/internal/kafka"
	"github.com/dm4brl/aut7emu/internal/mqtt"
	"github.com/dm4brl/aut7emu/internal/redis"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Println("Нет файла .env, используем стандартные настройки")
	}

	// Получаем параметры подключения из переменных окружения
mqttBroker := getEnv("MQTT_BROKER", "tcp://localhost:1883")
redisAddr := getEnv("REDIS_HOST", "localhost:6379")
kafkaBroker := getEnv("KAFKA_BROKER", "localhost:9092")
postgresURL := getEnv("POSTGRES_URL", "postgres://user:password@localhost:5432/device_db?sslmode=disable")


	deviceCount, err := strconv.Atoi(getEnv("DEVICE_COUNT", "1000"))
	if err != nil {
		deviceCount = 1000
	}

	// Инициализация всех сервисов
	mqttClient := mqtt.Connect(mqttBroker, "aut7emu-switch-emulator")
	if mqttClient == nil {
		log.Fatalf("Ошибка подключения к MQTT: %v", err)
	}
	log.Println("MQTT подключение установлено")

	redisClient := redis.Connect(redisAddr)
	if redisClient == nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	log.Println("Redis подключение установлено")

	kafkaWriter := kafka.Connect(kafkaBroker, "device-events")
	if kafkaWriter == nil {
		log.Fatalf("Ошибка подключения к Kafka: %v", err)
	}
	log.Println("Kafka подключение установлено")

	dbConn := db.Connect(postgresURL)
	if dbConn == nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}
	log.Println("Подключение к PostgreSQL установлено")

	// Эмуляция устройств
	for i := 1; i <= deviceCount; i++ {
		go devices.SimulateDevice(fmt.Sprintf("switch%d", i))
	}

	// Бесконечный цикл для предотвращения завершения программы
	select {}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

