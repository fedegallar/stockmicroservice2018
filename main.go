package main

import (
	"github.com/fedegallar/stockmicroservice2018/stock"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	f, _ := os.Create(time.currentTime.String()+".log")
	gin.DefaultWriter = io.MultiWriter(f)
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
	*     "5b97cd224a1aa37430480268" : 50,
	* }
	* @apiSuccessExample {json} Success-No-Content:
	* {
	*     "Code" : 204,
	*	  "Description" : "No content"
	* }
	*
	 */
	r.GET("/api/v1/stock/:articleid", stock.GetStockByArticleID)
	/**
	*
	* @api {POST} /api/v1/stock Add stock to an article. If there is no article, it will create a new one.
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
	r.POST("/api/v1/stock", stock.AddStockToArticle)
	/**
	*
	* @api {DELETE} /api/v1/stock/articleid Remove stock from an article.
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
	r.DELETE("api/v1/stock/:articleid", stock.RemoveStockFromArticle)
	r.Static("apidoc/", "./apidoc")
	r.Run(":3000")
}
