package main

import (
    "fmt"
    "log"
    "switch-emulator/internal/mqtt"
    "switch-emulator/internal/redis"
    "switch-emulator/internal/kafka"
    "switch-emulator/internal/db"
    "switch-emulator/internal/devices"
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

    // Бесконечный цикл
    select {}
}
