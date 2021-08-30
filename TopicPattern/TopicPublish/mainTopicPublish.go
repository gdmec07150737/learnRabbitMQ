package main

import (
	"fmt"
	"learnRabbitMQ/TopicPattern/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	pgcOne := RabbitMQ.NewRabbitMQTopic("exchangePgcTopic", "pgc.topic.one")
	pgcTwo := RabbitMQ.NewRabbitMQTopic("exchangePgcTopic", "pgc.topic.two")
	for i := 0; i <= 100; i++ {
		pgcOne.PublishTopic("Hello pgc topic one!" + strconv.Itoa(i))
		pgcTwo.PublishTopic("Hello pgc topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
