package producer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/md-mudassir7/go-kafka/config"
)

type Message struct {
	Text string `json:"text"`
}

var (
	producer sarama.SyncProducer
	cfg      config.Config
)

func init() {
	cfg = config.LoadConfig()

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = true

	brokers := []string{cfg.Kafka.Host + ":" + cfg.Kafka.Port}
	p, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	producer = p
}

func Publish(w http.ResponseWriter, r *http.Request) {
	message := &Message{}
	input_json, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(input_json, message)
	if err != nil {
		http.Error(w, "Error while unmarshaling request body", http.StatusBadRequest)
		return
	}
	messageToPublish := &sarama.ProducerMessage{
		Topic: cfg.Kafka.Topic,
		Value: sarama.StringEncoder(message.Text),
	}

	_, _, err = producer.SendMessage(messageToPublish)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		http.Error(w, "Error sending message", http.StatusInternalServerError)
	}
	log.Printf("Message published successfully on topic %s\n", cfg.Kafka.Topic)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message published successfully"))
}
