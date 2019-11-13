package main

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "os"
    "strings"
)

func bodyFrom(args []string) string {
    var s string
    if (len(args) < 2) || os.Args[1] == "" {
        s = "hello"
    } else {
        s = strings.Join(args[1:], " ")
    }
    return s
}

const AMQP_URL = "amqp://guest:guest@localhost:5672/"

func main() {

    // 接到RabbitMQ服务器(协议,身份认证)
    conn, err := amqp.Dial(AMQP_URL)
    defer conn.Close()
    if err != nil {
        fmt.Errorf("failed to connect to RabbitMQ: %s", err)
    }

    fmt.Println("amqp conn success ")

    // 创建一个通道
    ch, err := conn.Channel()
    if err != nil {
        fmt.Errorf("failed to open a channel: %s", err)
    }
    fmt.Println("conn channel success ")

    // 声明队列
    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    if err != nil {
        fmt.Errorf("failed to declare a queue: %s", err)
    }

    body := bodyFrom(os.Args)
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        })
    if err != nil {
        fmt.Errorf("%s : %s", "Failed to publish a message", err)
    }
    log.Printf(" [x] Sent %s", body)
}
