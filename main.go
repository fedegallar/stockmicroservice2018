package main

import (
	"time"
	"github.com/fedegallar/stockmicroservice2018/stock"
	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	/**
	*
	* @api {get} /stock/:articleid Request stock from an specific article.
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
	*     "5b97cd224a1aa37430480268" : 50,
	* }
	* @apiSuccessExample {json} Success-No-Content:
	* {
	*     "Code" : 204,
	*	  "Description" : "No content"
	* }
	*
	 */
	r.GET("/stock/:articleid", stock.GetStockByArticleID)
	/**
	*
	* @api {POST} /stock Add stock to an article. If there is no article, it will create a new one.
	* @apiName AddStockToAnArticle
	* @apiGroup Stock Operations
	* @apiVersion  1.0.0
	*
	* @apiParam {String} articleid The unique id of an article.
	* @apiParam {Number} quantity The quantity of an article.
	*
	* @apiSuccess (200) {String} articleid Id of the article.
	* @apiSuccess (200) {Number} quantity Number of articles aviable.
	*
	* @apiSuccessExample {json} Success:
	* {
	*     "5b97cd224a1aa37430480268" : 50,
	* }
	* @apiSuccessExample {json} Success-No-Content:
	* {
	*     "Code" : 204,
	*	  "Description" : "No content"
	* }
	*
	 */
	r.POST("/stock", stock.AddStockToArticle)
	/**
	*
	* @api {DELETE} /stock/articleid Remove stock from an article.
	* @apiName RemoveStockFromArticle
	* @apiGroup Stock Operations
	* @apiVersion  1.0.0
	*
	* @apiParam {String} articleid The unique id of an article.
	* @apiParam {Number} quantity The quantity of an article.
	*
	* @apiSuccess (200) {String} articleid Id of the article.
	* @apiSuccess (200) {Number} quantity Number of articles aviable.
	*
	* @apiSuccessExample {json} Success:
	* {
	*     "5b97cd224a1aa37430480268" : 50,
	* }
	* @apiSuccessExample {json} Success-No-Content:
	* {
	*     "Code" : 204,
	*	  "Description" : "No content"
	* }
	*
	 */
	r.DELETE("/stock/:articleid", stock.RemoveStockFromArticle)
	r.Static("apidoc/", "./apidoc")
	r.Run(":3000")
}
