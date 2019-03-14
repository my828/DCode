package handlers

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// RabbitStore allows to interact with the rabbit message queue
type RabbitStore struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

// NewRabbitStore initializes and returns a new rabbit store
func NewRabbitStore(mqAddress string, mqName string) *RabbitStore {
	connection, err := amqp.Dial(mqAddress)
	if err != nil {
		log.Println(err)
		return nil
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &RabbitStore{
		Connection: connection,
		Channel:    channel,
		QueueName:  mqName,
	}
}

// Consume reads from the rabbit message queue and returns a go channel
func (rs *RabbitStore) Consume() <-chan amqp.Delivery {
	queue, err := rs.Channel.QueueDeclare(
		rs.QueueName, // name
		false,        // durable
		false,        // delete when usused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Printf("error declaring a queue: %v\n", err)
	}
	messages, err := rs.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Printf("error consuming from queue: %v\n", err)
	}
	return messages
}

// Publish adds messages to the Rabbit message queue
func (rs *RabbitStore) Publish(message *Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		log.Println("error encoding json")
	}
	queue, err := rs.Channel.QueueDeclare(
		rs.QueueName, // name
		false,        // durable
		false,        // delete when usused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Printf("error declaring a queue: %v\n", err)
	}
	err = rs.Channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	return nil
}
