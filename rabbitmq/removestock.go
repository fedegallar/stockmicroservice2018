package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

//Article Es el articulo que Rabbitmq va autilizar
type Article struct {
	Articleid string `json:"articleid"`
	Quantity  int    `json:"quantity"`
}

func RemoveStock() {
	go func() {
		for {
			RemoveStockListener()
			time.Sleep(5 * time.Second)
		}
	}()
}

//RemoveStockListener Se queda escuchando para ver si es necesario remover stock.
func RemoveStockListener() {
	conn, err := amqp.Dial("amqp://localhost")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer chn.Close()
	err = chn.ExchangeDeclare(
		"sell_flow", // name
		"topic",     // type
		false,       // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		panic(err)
	}

	queue, err := chn.QueueDeclare(
		"Stock", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"order_placed", // routing key
		"sell_flow",    // exchange
		false,
		nil)
	if err != nil {
		panic(err)
	}
	mgs, err := chn.Consume(
		queue.Name,      // queue
		"StockConsumer", // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("RabbitMQ: A la espera para remover stock")
	forever := make(chan bool)
	go func() {
		for d := range mgs {
			log.Println("Article recived")
			articleContent := &Article{}
			err = json.Unmarshal(d.Body, articleContent)
			if err != nil {
				panic(err)
			}
			fmt.Println(articleContent.Articleid)
		}
	}()
	<-forever
}
