package main

import (
	"fmt"

	"github.com/hongry18/sarama-kafka-go/internal/consumer"
)

func main() {
	fmt.Println("Kafka Consumer Example")
	consumer.Consumer()
}
