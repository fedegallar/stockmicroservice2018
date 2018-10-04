package stock

import (
	"fmt"

	"github.com/fedegallar/stockmicroservice2018/article"
	"github.com/fedegallar/stockmicroservice2018/redisclient"
	"github.com/gin-gonic/gin"
)

//GetStockByArticleID Devuelve el stock de un articulo.
func GetStockByArticleID(c *gin.Context) {
	articleid := c.Param("articleid")
	fmt.Println(articleid)
	quantity := redisclient.GetStock(c)
	var result struct {
		articleid string
		quantity  string
	}
	fmt.Println(quantity)
	result.articleid = articleid
	result.quantity = quantity
	c.JSON(200, result)
}

//AddStockToArticle Agrega stock a un árticulo. Si no existe, lo crea.
func AddStockToArticle(c *gin.Context) {
	type Article struct {
		Articleid string `json:"articleid" binding:"required"`
		Quantity  string `json:"quantity" binding:"required"`
	}
	body := Article{}
	if err := c.ShouldBindJSON(&body); err != nil {
		panic(err)
	}
	article := article.New()
	article.Articleid = body.Articleid
	article.Quantity = body.Quantity
	redisclient.AddStock(article.Articleid, article.Quantity)
	c.JSON(200, body)
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
