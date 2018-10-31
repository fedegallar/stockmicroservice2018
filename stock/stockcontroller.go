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
	quantity, err := redisclient.GetStock(articleid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Article not found",
		})
		return
	}
	fmt.Println(quantity)
	c.JSON(200, gin.H{
		"articleid": articleid,
		"quantity":  quantity,
	})
}

//AddStockToArticle Agrega stock a un árticulo. Si no existe, lo crea.
func AddStockToArticle(c *gin.Context) {
	//Pregunta aca si ha iniciado sesión o se hace en main.
	type Article struct {
		Articleid string `json:"articleid" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}
	body := Article{}
	if err := c.ShouldBindJSON(&body); err != nil {
		panic(err)
	}
	article := article.New()
	article.Articleid = body.Articleid
	article.Quantity = body.Quantity
	msg, err := redisclient.ModifyStock(article.Articleid, article.Quantity)
	if err != nil {
		c.JSON(404, err)
	}
	c.JSON(200, msg)
}
