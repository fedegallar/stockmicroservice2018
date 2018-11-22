package redisclient

import (
	"fmt"
	"strconv"

	alert "github.com/fedegallar/stockmicroservice2018/rabbitmq/alert"

	"github.com/go-redis/redis"
)

func client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

//GetStock Trae el stock de un determinado artículo.
func GetStock(articleid string) (string, error) {
	client := client()
	val, err := client.Get(articleid).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

//AddStock Crea un articulo con stock.
func AddStock(articleid string, quantity int) (string, error) {
	err := client().Set(articleid, quantity, 0).Err()
	if err != nil {
		return "", err
	}
	return "Stock added successfully!", nil
}

//RemoveStock Deletes stock
func RemoveStock(articleid string, quantity int) (string, error) {
	val, err := client().Get(articleid).Result()
	if err == redis.Nil {
		fmt.Println("Article id doesn't exists:", articleid)
		return "", nil
	}
	var actualVal int
	actualVal, err = strconv.Atoi(val)
	quantity = quantity + actualVal
	if quantity < 10 {
		alert.LowStockAlert(articleid)
	}
	errset := client().Set(articleid, quantity, 0).Err()
	if errset != nil {
		fmt.Println("There was a problem removing stock:", articleid)
		return "", nil
	}
	return "", nil

}

//ModifyStock añade stock a un articulo. Si no existe, lo crea directamente.
func ModifyStock(articleid string, quantity int) (string, error) {
	val, err := client().Get(articleid).Result()
	if err == redis.Nil && val == "" {
		msg, errload := AddStock(articleid, quantity)
		if errload != nil {
			return "", errload
		}
		return msg, nil
	}
	errset := client().Set(articleid, quantity, 0).Err()
	if errset != nil {
		return "", errset
	}
	return "Stock added succesfully", nil
}
