package main

import (
	"fmt"

	"github.com/hongry18/sarama-kafka-go/internal/producer"
)

func main() {
	fmt.Println("Kafka Producer Example")
	producer.Producer()
}
