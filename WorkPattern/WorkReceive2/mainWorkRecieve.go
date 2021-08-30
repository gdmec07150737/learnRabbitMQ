package main

import "learnRabbitMQ/WorkPattern/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "pgcWorkQueue")
	rabbitmq.ConsumeSimple()
}