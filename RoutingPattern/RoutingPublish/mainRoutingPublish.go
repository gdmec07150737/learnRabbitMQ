package main

import (
	"fmt"
	"learnRabbitMQ/RoutingPattern/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	pgcOne := RabbitMQ.NewRabbitMQRouting("exchangePgc", "exchangePgcOne")
	pgcTwo := RabbitMQ.NewRabbitMQRouting("exchangePgc", "exchangePgcTwo")
	for i := 0; i <= 100; i++ {
		pgcOne.PublishRouting("Hello pgc one!" + strconv.Itoa(i))
		pgcTwo.PublishRouting("Hello pgc Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}