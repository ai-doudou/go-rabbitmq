package main

import (
    "fmt"
    "github.com/streadway/amqp"
    "time"
)

const AMQP_URL = "amqp://guest:guest@localhost:5672/"

// 发送消息
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

    // 发送消息体
    // 消息内容是一个字节数组，因此您可以在此处编码任何内容。
    body := "Hello World! ["+time.Now().Format("2006-01-02 15:04:05")+"]"
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    if err != nil {
        fmt.Errorf("failed to publish a message: %s", err)
    }

}
