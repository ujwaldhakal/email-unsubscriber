package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func getConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(queueName string, message []byte) {
	connection := getConnection()
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        message,
		})
	failOnError(err, "Failed to publish a messages")
	defer ch.Close()
}

func ConsumerClient(queueName string) <-chan amqp.Delivery {
	connection := getConnection()
	ch, _ := connection.Channel()

	msgs, _ := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	return msgs
}
