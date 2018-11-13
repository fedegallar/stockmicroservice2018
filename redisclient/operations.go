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
			return "", errload //DEBERIA DEVOLVER UN 500 PORQUE SE CHINGO LA BASE DE DATOS!
		}
		return msg, nil
	}
	var actualVal int
	actualVal, err = strconv.Atoi(val)
	if err != nil {
		return "", err //VER AL ULTIMO COMO TIENE QUE DEVOLVER UN 400 BAD REQUEST!!!
	}
	quantity = quantity + actualVal
	errset := client().Set(articleid, quantity, 0).Err()
	if errset != nil {
		return "", errset //DEBERIA DEVOLVER 500 TAMBIEN!!!
	}
	return "Added succesfully", nil
}
