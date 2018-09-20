package main

import (
	"github.com/fedegallar/stockmicroservice2018/stock"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/**
	*
	* @api {get} /stock Request all Stock aviable
	* @apiName GetAllStock
	* @apiGroup Stock Information
	* @apiVersion  1.0.0
	*
	* @apiSuccess (200) {String} articleid Id of the article.
	* @apiSuccess (200) {Number} quantity Number of articles aviable.
	*
	* @apiSuccessExample {json} Success:
	* {
	*     "5b97cd224a1aa37430480268" : 50,
	*	   "5b97afdal1209u7cjm564735" : 20
	* }
	* @apiSuccessExample {json} Success-No-Content:
	* {
	*     "Code" : 204,
	*	  "Description" : "No content"
	* }
	*
	 */
	r.GET("/stock", stock.GetAllArticles)
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
	r.GET("/stock/:articleid/:quantity", stock.GetStockByArticleId)
	/**
	*
	* @api {put} /stock/:articleid/:quantity Add new article with stock.
	* @apiName AddNewArticleWithStock
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
	*
	 */
	r.PUT("/stock/:articleid/:quantity", stock.AddNewArticle)
	/**
	*
	* @api {POST} /stock/add/:articleid/:quantity Add stock to an article.
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
	r.POST("/stock/add/:articleid/:quantity", stock.AddStockToArticle)
	/**
	*
	* @api {POST} /stock/remove/:articleid/:quantity Remove stock from an article.
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
	r.POST("/stock/remove/:articleid/:quantity", stock.RemoveStockFromArticle)
	r.Static("apidoc/", "./apidoc")
	r.Run(":3000")
}
