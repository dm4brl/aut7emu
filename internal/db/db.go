package db

import (
    "fmt"
    "log"
    "github.com/jackc/pgx/v4"
    "context"
)

var db *pgx.Conn
var ctx = context.Background()

// Инициализация подключения к PostgreSQL
func Connect(postgresURL string) *pgx.Conn {
    var err error
    db, err = pgx.Connect(ctx, postgresURL)
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных: ", err)
    }

    fmt.Println("Подключено к PostgreSQL")
    return db
}

// Сохранение состояния устройства в базу данных
func SaveDeviceState(deviceID string, state string) {
    _, err := db.Exec(ctx, "INSERT INTO device_states (device_id, state) VALUES ($1, $2)", deviceID, state)
    if err != nil {
        log.Printf("Ошибка сохранения состояния устройства в базу данных: %v", err)
    }
}
