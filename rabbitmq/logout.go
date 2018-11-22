package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fedegallar/stockmicroservice2018/config/errors"
	"github.com/fedegallar/stockmicroservice2018/security"
	"github.com/streadway/amqp"
)

type message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Init se queda escuchando broadcasts de logout
func Init() {
	go func() {
		for {
			listenLogout()
			fmt.Println("RabbitMQ conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

/**
 * @api {fanout} auth/logout User logout
 * @apiGroup RabbitMQ GET
 *
 * @apiDescription Listen logout messages from auth.
 *
 * @apiSuccessExample {json} Message
 *     {
 *        "type": "logout",
 *        "message": "{tokenId}"
 *     }
 */
func listenLogout() error {
	conn, err := amqp.Dial("amqp://localhost")
	if err != nil {
		println(errors.ConnectionError)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		println(errors.ChannelError)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"auth",   // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		println(errors.ExchangeDeclareError)
		return err
	}

	queue, err := chn.QueueDeclare(
		"auth", // name
		false,  // durable
		false,  // delete when unused
		true,   // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		println(errors.QueueDeclareError)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
		println(errors.QueueBindError)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		println(errors.ConsumeError)
		return err
	}

	fmt.Println("RabbitMQ: Logout")

	go func() {
		for d := range mgs {
			log.Output(1, "Mensage recibido")
			newMessage := &message{}
			err = json.Unmarshal(d.Body, newMessage)
			if err == nil {
				if newMessage.Type == "logout" {
					security.Invalidate(newMessage.Message)
				}
			}
		}
	}()

	fmt.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}
