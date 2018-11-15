package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

//LowStockAlert Emite alerta de stock bajo
func LowStockAlert(articleid string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		println("Error al abrir una conexion con RabbitMQ\n", err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		println("Error al abrir un Channel\n", err)
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
		println("Error al publicar un mensaje", err)
	}
}
