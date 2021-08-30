package main

import "learnRabbitMQ/RoutingPattern/RabbitMQ"

func main() {
	pgcOne := RabbitMQ.NewRabbitMQRouting("exchangePgc", "exchangePgcOne")
	pgcOne.ReceiveRouting()
}