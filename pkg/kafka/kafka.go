package kafka

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type Messaging interface {
	SendMessage(urlImage string, typeDetect string) (string, error)
	ChannelMessage(ctx context.Context, url string, typeDetect string) (bool, string, error)
	ConsumerMessage(messageKey string) (string, error)
	SendAndReceiver(ctx context.Context, urlImage, typeDetect string) (string, error)
	SendAndReceiverService(ctx context.Context, urlImage, typeDetect string) (string, error)
}

type Kafka struct {
}

func NewKafka() Messaging {
	return &Kafka{}
}

func (k *Kafka) SendMessage(urlImage string, typeDetect string) (string, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"kafka:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("erro ao criar produtor %w", err)
	}

	defer producer.Close()

	requestID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("erro ao gerar id da request %w", err)
	}

	message := &sarama.ProducerMessage{
		Topic: "pending-analysis",
		Key:   sarama.StringEncoder(requestID.String()),
		Value: sarama.StringEncoder(urlImage),
		Headers: []sarama.RecordHeader{
			{Key: []byte("key"), Value: []byte(typeDetect)},
		},
	}

	partiton, offset, err := producer.SendMessage(message)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar mensagem %w", err)
	}

	fmt.Println(partiton, offset)

	return requestID.String(), nil
}

func (k *Kafka) ChannelMessage(ctx context.Context, url string, typeDetect string) (bool, string, error) {

	type KafkaSendMessage struct {
		MessageKey string
		Success    bool
		Err        error
	}

	resultSendedMessage := make(chan KafkaSendMessage)

	go func() {

		messageKey, err := k.SendMessage(url, typeDetect)
		resultSendedMessage <- KafkaSendMessage{Success: err == nil, Err: err, MessageKey: messageKey}

	}()

	select {

	case <-ctx.Done():
		return false, "", ctx.Err()

	case result := <-resultSendedMessage:
		return result.Success, result.MessageKey, result.Err
	}
}

func (k *Kafka) ConsumerMessage(messageKey string) (string, error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"kafka:9092"}

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		return "", fmt.Errorf("erro ao criar consumer: %w", err)
	}

	defer consumer.Close()

	topicName := "result-analysis"

	partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)

	if err != nil {
		return "", fmt.Errorf("erro ao definir particao: %w", err)
	}

	for {
		select {
		case result := <-partitionConsumer.Messages():
			if string(result.Key) == (messageKey) {
				return string(result.Value), nil
			}
		case err := <-partitionConsumer.Errors():
			return "", err
		case <-time.After(5 * time.Second):
			return "", fmt.Errorf("timeout: mensagem com key %s não encontrada", messageKey)
		}
	}
}

func (k *Kafka) SendAndReceiver(ctx context.Context, urlImage, typeDetect string) (string, error) {

	successSendedChannelMessage, messageKey, err := k.ChannelMessage(ctx, urlImage, typeDetect)

	if !successSendedChannelMessage {
		log.Println(err)

		if errors.Is(err, ctx.Err()) {
			return "", fmt.Errorf("%s", myerror.ErrTimeOut)
		}

		return "", fmt.Errorf("%s", err.Error())
	}

	message, err := k.ConsumerMessage(messageKey)
	if err != nil {

		switch errTyped := err.(type) {
		case *sarama.ConsumerError:
			return "", errTyped

		default:
			return "", errTyped
		}

	}

	return message, nil
}

func (k *Kafka) SendAndReceiverService(ctx context.Context, urlImage, typeDetect string) (string, error) {

	message, err := k.SendAndReceiver(ctx, urlImage, typeDetect)

	if err != nil {

		if errors.Is(err, ctx.Err()) {
			log.Println(err.Error())
			return "", fmt.Errorf("tempo excedido ao tentar se comunicar com o kafka")
		}

		if errors.As(err, &sarama.ConsumerError{}) {
			log.Println(err.Error())
			return "", fmt.Errorf("erro ao consumir mensagem do kafka")
		}

		log.Println(err.Error())
		return "", fmt.Errorf("erro ao se comunicar com o servidor")
	}

	return message, nil
}
