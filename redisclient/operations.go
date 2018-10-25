package redisclient

import (
	"strconv"

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

//ModifyStock añade stock a un articulo. Si no existe, lo crea directamente.
func ModifyStock(articleid string, quantity int) (string, error) {
	val, err := client().Get(articleid).Result()
	if err == redis.Nil {
		msg, errload := AddStock(articleid, quantity)
		if errload != nil {
			return "", errload
		}
		return msg, nil
	} else {
		var actualVal int
		actualVal, err = strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		quantity = quantity + actualVal
		errset := client().Set(articleid, quantity, 0).Err()
		if errset != nil {
			return "", errset
		}
		return "Added succesfully", nil
	}
}
