package main

import (
    "github.com/streadway/amqp"
    "log"
)

const AMQP_URL = "amqp://guest:guest@localhost:5672/"

func main() {

    conn, err := amqp.Dial(AMQP_URL)
    if err != nil {
        log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("%s: %s", "Failed to open a channe", err)
    }

    // 声明队列
    q, err := ch.QueueDeclare(
        "hello", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        log.Fatalf("%s: %s", "Failed to register a consumer", err)
    }

    forever := make(chan bool)
    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever

}
