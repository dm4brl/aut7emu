package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dm4brl/aut7emu/internal/db"
	"github.com/dm4brl/aut7emu/internal/devices"
	"github.com/dm4brl/aut7emu/internal/mqtt"
	"github.com/dm4brl/aut7emu/internal/redis"
	"github.com/dm4brl/aut7emu/internal/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
    // Инициализация всех сервисов
    mqtt.Connect("tcp://localhost:1883", "switch-emulator")
    redis.Connect("localhost:6379")
    kafka.Connect("localhost:9092", "device-events")
    db.Connect("postgres://user:password@localhost:5432/device_db")

    // Эмуляция нескольких устройств
    for i := 1; i <= 1000; i++ {
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

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}