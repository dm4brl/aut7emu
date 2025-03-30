package mqtt

import (
    "fmt"
    "log"
    "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

// Инициализация подключения к EMQX
func Connect(broker string, clientID string) mqtt.Client {
    opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
    opts.OnConnect = func(c mqtt.Client) {
        fmt.Println("Подключено к EMQX")
    }

    client = mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    return client
}

// Публикация сообщения на канал
func Publish(topic string, payload string) {
    if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
        log.Printf("Ошибка публикации: %v", token.Error())
    }
}
