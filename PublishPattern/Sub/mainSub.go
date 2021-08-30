package main

import "learnRabbitMQ/PublishPattern/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newExchange")
	rabbitmq.ReceiveSub()
}