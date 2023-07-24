package consumer

import (
	"fmt"

	"github.com/IBM/sarama"
)

func Consumer() {
	var nTndIdx int32 = 0
	config := sarama.NewConfig()

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	lastOffset, err := client.GetOffset("topic-1", nTndIdx, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("topic-1", nTndIdx, lastOffset)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	consumed := 0
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Topic %s Consumed message offset %d, Partition %d\n", msg.Topic, msg.Offset, msg.Partition)
			consumed++
			fmt.Printf("Consumed: %d\n", consumed)
			fmt.Println(string(msg.Key))
			fmt.Println(string(msg.Value))
			fmt.Println("")
		}
	}
}
