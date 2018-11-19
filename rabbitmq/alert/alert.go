package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

/**
*
* @api {topic} stock_alert Low stock alert.
* @apiGroup RabbitMQ POST
* @apiVersion  1.0.0
* @apiDescription Emits an low stock alert.
*
* @apiParam {String} articleid The unique id of an article.
*
* @apiSuccessExample {text/plain} Mensaje
* articleid
*
*
 */
func LowStockAlert(articleid string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		println("There was an error opening a connection with RabbitMQ\n", err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		println("There was an error opening a channel\n", err)
		return
	}
	defer ch.Close()
	fmt.Println("Lowstock alert for: ", articleid)
	err = ch.Publish(
		"stock_alert",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(articleid),
		})
	if err != nil {
		println("There was an error publishing a message", err)
	}
}
