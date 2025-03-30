# aut7emu

├── cmd/               # Основной код сервера эмулятора
├── internal/          # Логика работы с MQTT, Redis, Kafka и базой данных
│   ├── mqtt/          # Подключение к EMQX
│   ├── redis/         # Взаимодействие с Redis
│   ├── kafka/         # Работа с Kafka
│   ├── db/            # Взаимодействие с PostgreSQL
│   └── devices/       # Модели и логика устройства
├── configs/           # Конфигурационные файлы
├── Dockerfile         # Docker файл для сборки контейнера
├── README.md          # Документация проекта
