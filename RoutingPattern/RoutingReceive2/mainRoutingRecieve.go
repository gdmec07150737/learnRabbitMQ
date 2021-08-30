package main

import "learnRabbitMQ/RoutingPattern/RabbitMQ"

func main() {
	pgcTwo := RabbitMQ.NewRabbitMQRouting("exchangePgc", "exchangePgcTwo")
	pgcTwo.ReceiveRouting()
}