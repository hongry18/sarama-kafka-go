package producer

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/hongry18/sarama-kafka-go/model"
)

func Producer() {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	connectionString := []string{"127.0.0.1:9092"}

	conn, err := sarama.NewClient(connectionString, conf)
	if err != nil {
		panic(err)
	}

	producer, err := sarama.NewSyncProducerFromClient(conn)
	if err != nil {
		panic(err)
	}

	data := model.Sample{ID: 5, Name: "test-1"}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	msg := &sarama.ProducerMessage{
		Topic:     "topic-1",
		Key:       sarama.StringEncoder("id-1"),
		Value:     sarama.ByteEncoder(jsonBody),
		Partition: int32(0),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
}
