package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

func SendMessageKafka(urlImage string) error {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"kafka:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("erro ao criar produtor")
	}

	defer producer.Close()

	requestID, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	message := &sarama.ProducerMessage{
		Topic: "pending-analysis",
		Key:   sarama.StringEncoder(requestID.String()),
		Value: sarama.StringEncoder(urlImage),
	}

	partiton, offset, err := producer.SendMessage(message)
	if err != nil {
		panic(err)
	}

	fmt.Println(partiton, offset)

	return nil
}

func KafkaChannelMessage(ctx context.Context, url string) (bool, error) {

	type KafkaSendMessage struct {
		Success bool
		Err     error
	}

	resultSendedMessage := make(chan KafkaSendMessage)

	go func() {

		err := SendMessageKafka(url)
		resultSendedMessage <- KafkaSendMessage{Success: err == nil, Err: err}

	}()

	select {

	case <-ctx.Done():
		return false, ctx.Err()

	case result := <-resultSendedMessage:
		return result.Success, result.Err
	}
}
