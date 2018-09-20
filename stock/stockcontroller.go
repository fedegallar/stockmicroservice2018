package stock

import (
	"github.com/gin-gonic/gin"
)

func GetStockByArticleId(c *gin.Context) {
	articleid := c.Param("articleid")
	c.JSON(200, gin.H{
		articleid: "Stock de articulo encontrado",
	})
}

func GetAllArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Todos los articulos van por aqui!",
	})
}

func AddNewArticle(c *gin.Context) {
	articleid := c.Param("articleid")
	quantity := c.Param("quantity")
	c.JSON(200, gin.H{
		"message":   "Article added succesfully",
		"articleid": articleid,
		"quantity":  quantity,
	})
}

func AddStockToArticle(c *gin.Context) {
	articleid := c.Param("articleid")
	quantity := c.Param("quantity")
	c.JSON(200, gin.H{
		"message":   "Stock added succesfully",
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
