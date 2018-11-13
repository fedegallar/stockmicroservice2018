package stock

import (
	"fmt"

	"github.com/fedegallar/stockmicroservice2018/article"
	"github.com/fedegallar/stockmicroservice2018/config/errors"
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
	err := validateAuthentication(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	type Article struct {
		Articleid string `json:"articleid" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}
	body := Article{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, errors.BadRequest)
		return
	}
	article := article.New()
	article.Articleid = body.Articleid
	article.Quantity = body.Quantity

	//REVISAR ESTA PORCION. ESTA HECHO POR SI NO LO ENCUENTRA... DIRECTAMENTE AÑADE UN ARTICULO NUEVO
	msg, err := redisclient.ModifyStock(article.Articleid, article.Quantity)
	if err != nil {
		c.JSON(500, errors.InternalError)
	}
	c.JSON(200, msg)
}
