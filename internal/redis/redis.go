package redis

import (
    "log"
    "github.com/go-redis/redis/v8"
    "context"
)

var client *redis.Client
var ctx = context.Background()

// Инициализация клиента Redis
func Connect(redisAddr string) *redis.Client {
    client = redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })

    _, err := client.Ping(ctx).Result()
    if err != nil {
        log.Fatal("Не удалось подключиться к Redis: ", err)
    }

    return client
}

// Сохранение состояния устройства
func SetDeviceState(deviceID string, state string) {
    err := client.Set(ctx, deviceID, state, 0).Err()
    if err != nil {
        log.Printf("Ошибка сохранения состояния устройства: %v", err)
    }
}

// Получение состояния устройства
func GetDeviceState(deviceID string) (string, error) {
    state, err := client.Get(ctx, deviceID).Result()
    if err != nil {
        return "", err
    }
    return state, nil
}
