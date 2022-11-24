package services

import (
	"encoding/json"
	"fmt"

	"github.com/Rickykn/emas-digital-api/database"
	"github.com/Rickykn/emas-digital-api/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReceiveFromKafka() {

	// fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"input-harga"}, nil)

	msg, err := c.ReadMessage(-1)

	if err == nil {
		job := string(msg.Value)
		go saveToDb(job)
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

	c.Close()

}

func saveToDb(jobString string) {

	DB := database.Get()

	var _price *models.Price
	b := []byte(jobString)
	err := json.Unmarshal(b, &_price)
	if err != nil {
		panic(err)
	}
	fmt.Println(_price)

	result := DB.Create(_price)
	if result.Error != nil {
		panic(result.Error)
	}
}
