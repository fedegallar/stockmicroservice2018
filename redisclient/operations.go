package redisclient

import (
	"github.com/gin-gonic/gin"
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
func GetStock(c *gin.Context) string {
	client := client()
	val, err := client.Get(c.Param("articleid")).Result()
	if err != nil {
		panic(err)
	}
	return val
}

//AddStock Añade stock a un articulo. Si el articulo no existe, crea uno nuevo.
func AddStock(articleid string, quantity string) {
	err := client().Set(articleid, quantity, 0).Err()
	if err != nil {
		panic(err)
	}
}
