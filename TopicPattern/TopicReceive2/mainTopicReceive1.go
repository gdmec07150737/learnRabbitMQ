package main

import "learnRabbitMQ/TopicPattern/RabbitMQ"

func main() {
	pgcOne := RabbitMQ.NewRabbitMQTopic("exchangePgcTopic", "pgc.*.two")
	pgcOne.ReceiveTopic()
}