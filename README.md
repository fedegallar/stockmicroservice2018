<a name="top"></a>
# Stock Microservice 2018 v1.0.0

This is an stock microservice made in Go

- [RabbitMQ_GET](#rabbitmq_get)
	- [User logout](#user-logout)
	- [Remove stock from an article.](#remove-stock-from-an-article.)
	
- [RabbitMQ_POST](#rabbitmq_post)
	- [Low stock alert.](#low-stock-alert.)
	
- [Stock_Information](#stock_information)
	- [Request stock from an specific article.](#request-stock-from-an-specific-article.)
	
- [Stock_Operations](#stock_operations)
	- [Add stock to an article.](#add-stock-to-an-article.)
	


# <a name='rabbitmq_get'></a> RabbitMQ_GET

## <a name='user-logout'></a> User logout
[Back to top](#top)

<p>Listen logout messages from auth.</p>

	FANOUT auth/logout





### Success Response

Message

```
{
   "type": "logout",
   "message": "{tokenId}"
}
```


## <a name='remove-stock-from-an-article.'></a> Remove stock from an article.
[Back to top](#top)

<p>Listen messages from order and remove stock when an order is placed.</p>

	TOPIC sell_flow/order_placed





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  articleid | String | <p>The unique id of an article.</p>|
|  quantity | Number | <p>The quantity of an article.</p>|




# <a name='rabbitmq_post'></a> RabbitMQ_POST

## <a name='low-stock-alert.'></a> Low stock alert.
[Back to top](#top)

<p>Emits an low stock alert.</p>

	TOPIC stock_alert





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  articleid | String | <p>The unique id of an article.</p>|


### Success Response

Mensaje

```
articleid
```


# <a name='stock_information'></a> Stock_Information

## <a name='request-stock-from-an-specific-article.'></a> Request stock from an specific article.
[Back to top](#top)



	GET /api/v1/stock/:articleid





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  articleid | Number | <p>The unique id of an article.</p>|


### Success Response

Success:

```
{
    "<articleid>" : <quantity>,
}
```

### 200

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  articleid | String | <p>Id of the article.</p>|
|  quantity | Number | <p>Number of articles aviable.</p>|

### Error Response

ArticleNotFound:

```

HTTP/1.1 404 Not found
{"error":"Article not found"}
```
# <a name='stock_operations'></a> Stock_Operations

## <a name='add-stock-to-an-article.'></a> Add stock to an article.
[Back to top](#top)



	POST /api/v1/stock/:articleid





### Parameter Parameters

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  articleid | String | <p>The unique id of an article.</p>|
|  quantity | Number | <p>The quantity of an article.</p>|


### Success Response

Success:

```
HTTP/1.1 200 Success
{
    "message":"Article added successfully!"
}
```


### Error Response

BadRequest:

```
HTTP/1.1 400 Bad Request
{
	  "error" : "Parameters needed"
}
```
NotAutorized:

```
HTTP/1.1 401 Not Autorized
{
	  "error" : "Not Autorized"
}
```
ArticleNotFound:

```
HTTP/1.1 404 Not Found
{
	  "error" : "Article not found"
}
```
