# Используем официальный образ Go
FROM golang:1.20-alpine AS build

# Установим зависимости для сборки
RUN apk add --no-cache git

# Рабочая директория
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Сборка бинарного файла
RUN go mod tidy
RUN go build -o switch-emulator .

# Создаем финальный образ
FROM alpine:latest

# Устанавливаем зависимости для выполнения (MQTT, Redis, Kafka)
RUN apk add --no-cache ca-certificates

# Копируем собранный бинарник
COPY --from=build /app/switch-emulator /usr/local/bin/switch-emulator

# Порт, который будет слушать приложение
EXPOSE 8080

# Запуск приложения
CMD ["/usr/local/bin/switch-emulator"]
