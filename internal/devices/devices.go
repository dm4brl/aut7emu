package devices

import (
    "fmt"
    "time"
    "math/rand"
    "github.com/dm4brl/aut7emu/internal/mqtt"
    "github.com/dm4brl/aut7emu/internal/redis"
    "github.com/dm4brl/aut7emu/internal/kafka"
    "github.com/dm4brl/aut7emu/internal/db"
)

// Эмуляция устройства, которое периодически меняет свое состояние
func SimulateDevice(deviceID) {
    states := []string{"ON", "OFF"}
    for {
        state := states[rand.Intn(len(states))] // случайное состояние

        // Публикуем в MQTT
        mqtt.Publish("home/lighting/" + deviceID + "/status", state)

        // Сохраняем состояние в Redis
        redis.SetDeviceState(deviceID, state)

        // Отправляем сообщение в Kafka
        kafka.SendMessage(kafka.Connect("localhost:9092", "device-events"), fmt.Sprintf("Device %s state: %s", deviceID, state))

        // Сохраняем в базе данных
        db.SaveDeviceState(deviceID, state)

        // Задержка в 15 минут
        time.Sleep(15 * time.Minute)
    }
}
