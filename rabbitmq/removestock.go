package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fedegallar/stockmicroservice2018/config/errors"
	"github.com/fedegallar/stockmicroservice2018/redisclient"
	"github.com/streadway/amqp"
)

/**
*
* @api {topic} sell_flow/order_placed Remove stock from an article.
* @apiGroup RabbitMQ GET
* @apiVersion  1.0.0
* @apiDescription Listen messages from order and remove stock when an order is placed.
*
* @apiParam {String} articleid The unique id of an article.
* @apiParam {Number} quantity The quantity of an article.
*
*
*
 */

//Article define un articulo
type Article struct {
	Articleid string `json:"articleid"`
	Quantity  int    `json:"quantity"`
}

//Message Mensaje con todos los articulos a descontar
type Message struct {
	Articles []Article `json:"articles"`
}

//Object Es el articulo que Rabbitmq va autilizar
type Object struct {
	Type string  `json:"type"`
	Msg  Message `json:"message"`
}

//RemoveStock Keep listening for a message that removes stock.
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
		println(errors.ConnectionError)
		panic(err)
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		println(errors.ChannelError)
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
		println(errors.ExchangeDeclareError)
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
		println(errors.QueueDeclareError)
		panic(err)
	}

	err = chn.QueueBind(
		queue.Name,     // queue name
		"order_placed", // routing key
		"sell_flow",    // exchange
		false,
		nil)
	if err != nil {
		println(errors.QueueBindError)
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
		println(errors.ConsumeError)
		panic(err)
	}

	fmt.Println("RabbitMQ: Remove Stock waiting...")
	forever := make(chan bool)
	go func() {
		for d := range mgs {
			newArticle := &Object{}
			log.Println("Article recived")
			err = json.Unmarshal(d.Body, newArticle)
			for _, art := range newArticle.Msg.Articles {
				fmt.Println("Getting data: ", art.Articleid, " ", art.Quantity*(-1))
				redisclient.RemoveStock(art.Articleid, art.Quantity*(-1))
			}
		}
	}()
	<-forever
}
