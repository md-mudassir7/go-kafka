package consumer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/md-mudassir7/go-kafka/config"
)

var (
	consumer sarama.Consumer
	cfg      config.Config
)

func init() {
	cfg = config.LoadConfig()
	brokers := []string{cfg.Kafka.Host + ":" + cfg.Kafka.Port}
	c, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	consumer = c
}

func ConsumeMessages() {
	partitionConsumer, err := consumer.ConsumePartition(cfg.Kafka.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				log.Printf("Received message: %s on topic %s\n", msg.Value, cfg.Kafka.Topic)
			// ctrl + c to close the consumer
			case <-signals:
				log.Println("Interrupt signal received, closing consumer...")
				doneCh <- struct{}{}
				return
			}
		}
	}()

	<-doneCh
}
