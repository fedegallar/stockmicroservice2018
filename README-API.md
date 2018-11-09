<a name="top"></a>
# Stock Microservice 2018 v1.0.0

This is an stock microservice made in Go

- [Stock_Information](#stock_information)
	- [Request stock from an specific article.](#request-stock-from-an-specific-article.)
	
- [Stock_Operations](#stock_operations)
	- [Add stock to an article.](#add-stock-to-an-article.)
	


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
