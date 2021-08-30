package main

import (
	"fmt"
	"learnRabbitMQ/SimplePattern/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "pgcQueue")
	rabbitmq.PublishSimple("Hello pgcQueue!")
	fmt.Println("发送成功！")
}