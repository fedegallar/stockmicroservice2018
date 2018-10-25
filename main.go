package main

import (
	"github.com/fedegallar/stockmicroservice2018/stock"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/**
	*
	* @api {get} /api/v1/stock/:articleid Request stock from an specific article.
	* @apiName GetStockFromAnArticle
	* @apiGroup Stock Information
	* @apiVersion  1.0.0
	*
	* @apiParam {Number} articleid The unique id of an article.
	*
	* @apiSuccess (200) {String} articleid Id of the article.
	* @apiSuccess (200) {Number} quantity Number of articles aviable.
	*
	* @apiSuccessExample {json} Success:
	* {
	*     "<articleid>" : <quantity>,
	* }
	*
	* @apiErrorExample ArticleNotFound:
	*
	* HTTP/1.1 404 Not found
	* {"error":"Article not found"}
	*
	*
	 */
	r.GET("/api/v1/stock/:articleid", stock.GetStockByArticleID)
	/**
	*
	* @api {POST} /api/v1/stock/:articleid Add stock to an article.
	* @apiName AddStockToAnArticle
	* @apiGroup Stock Operations
	* @apiVersion  1.0.0
	*
	* @apiParam {String} articleid The unique id of an article.
	* @apiParam {Number} quantity The quantity of an article.
	*
	*
	* @apiSuccessExample {json} Success:
	* HTTP/1.1 200 Success
	* {
	*     "message":"Article added successfully!"
	* }
	* @apiErrorExample {json} BadRequest:
	* HTTP/1.1 400 Bad Request
	* {
	*	  "error" : "Parameters needed"
	* }
	*
	* @apiErrorExample {json} NotAutorized:
	* HTTP/1.1 401 Not Autorized
	* {
	*	  "error" : "Not Autorized"
	* }
	*
	* @apiErrorExample {json} ArticleNotFound:
	* HTTP/1.1 404 Not Found
	* {
	*	  "error" : "Article not found"
	* }
	*
	 */
	r.POST("/api/v1/stock/:articleid", stock.AddStockToArticle)
	r.GET("/", func(c *gin.Context) {
		c.JSON(418, gin.H{
			"error": "I'm a teapot",
		})
	})
	r.Static("apidoc/", "./apidoc")
	r.Run(":3000")
}
