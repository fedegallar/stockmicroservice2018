package stock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/fedegallar/stockmicroservice2018/redisclient"
)
//GetStockByArticleID Devuelve el stock de un articulo.
func GetStockByArticleID(c *gin.Context) {
	articleid := c.Param("articleid")
	fmt.Println(articleid)
	quantity := redisclient.GetStock(c)
		var result struct{
			articleid string
			quantity string
		}
		result.articleid = articleid
		result.quantity = quantity
		c.JSON(200,result)
}

//AddStockToArticle Agrega stock a un árticulo. Si no existe, lo crea.
func AddStockToArticle(c *gin.Context) {
	articleid := c.Param("articleid")
	quantity := c.Param("quantity")
	result := redisclient.AddStock(articleid,quantity)
	c.JSON(200, gin.H{
		"message":   result,
		"articleid": articleid,
		"quantity":  quantity,
	})
}

func RemoveStockFromArticle(c *gin.Context) {
	articleid := c.Param("articleid")
	quantity := c.Param("quantity")
	c.JSON(200, gin.H{
		"message":   "Stock removed succesfully",
		"articleid": articleid,
		"quantity":  quantity,
	})
}
