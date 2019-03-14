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
	Queue      amqp.Queue
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
	queue, err := channel.QueueDeclare(
		mqName, // name
		false,  // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &RabbitStore{
		Connection: connection,
		Channel:    channel,
		Queue:      queue,
		QueueName:  mqName,
	}
}

// Consume reads from the rabbit message queue and returns a go channel
func (rs *RabbitStore) Consume() <-chan amqp.Delivery {
	messages, err := rs.Channel.Consume(
		rs.Queue.Name, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
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
	err = rs.Channel.Publish(
		"",
		rs.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	
	// for testing
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

