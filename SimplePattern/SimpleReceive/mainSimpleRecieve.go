package main

import "learnRabbitMQ/SimplePattern/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "pgcQueue")
	rabbitmq.ConsumeSimple()
}