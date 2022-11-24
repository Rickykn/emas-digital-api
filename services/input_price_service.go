package services

import (
	"encoding/json"
	"net/http"

	"github.com/Rickykn/emas-digital-api/helpers"
	"github.com/Rickykn/emas-digital-api/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InputPrice(w http.ResponseWriter, r *http.Request) *models.PriceDTO {
	body := &models.PriceDTO{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return nil
	}

	saveToKafka(body)

	return body
}

func saveToKafka(price *models.PriceDTO) {

	jsonString, _ := json.Marshal(price)

	jobString := string(jsonString)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	topic := "input-harga"
	for _, word := range []string{string(jobString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
}
