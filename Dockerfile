# Используем официальный образ Go
FROM golang:1.20-alpine AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum в контейнер (для кеширования зависимостей)
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь код проекта в контейнер
COPY . .

# Строим проект
RUN go build -o switch-emulator ./cmd

# Создаем финальный образ
FROM alpine:latest

# Устанавливаем зависимости для выполнения (опционально)
RUN apk add --no-cache ca-certificates

# Копируем собранный бинарник из предыдущего этапа
COPY --from=build /app/switch-emulator /usr/local/bin/switch-emulator

# Стартуем приложение
CMD ["/usr/local/bin/switch-emulator"]