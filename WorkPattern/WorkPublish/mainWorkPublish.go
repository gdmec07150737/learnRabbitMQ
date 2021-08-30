package main

import (
	"fmt"
	"learnRabbitMQ/WorkPattern/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "pgcWorkQueue")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello pgcWorkQueue!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}