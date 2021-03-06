define({ "api": [
  {
    "type": "fanout",
    "url": "auth/logout",
    "title": "User logout",
    "group": "RabbitMQ_GET",
    "description": "<p>Listen logout messages from auth.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "{\n   \"type\": \"logout\",\n   \"message\": \"{tokenId}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbitmq/logout.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "FanoutAuthLogout"
  },
  {
    "type": "topic",
    "url": "sell_flow/order_placed",
    "title": "Remove stock from an article.",
    "group": "RabbitMQ_GET",
    "version": "1.0.0",
    "description": "<p>Listen messages from order and remove stock when an order is placed.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "articleid",
            "description": "<p>The unique id of an article.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "quantity",
            "description": "<p>The quantity of an article.</p>"
          }
        ]
      }
    },
    "filename": "./rabbitmq/removestock.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "TopicSell_flowOrder_placed"
  },
  {
    "type": "topic",
    "url": "stock_alert",
    "title": "Low stock alert.",
    "group": "RabbitMQ_POST",
    "version": "1.0.0",
    "description": "<p>Emits an low stock alert.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "articleid",
            "description": "<p>The unique id of an article.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Mensaje",
          "content": "articleid",
          "type": "text/plain"
        }
      ]
    },
    "filename": "./rabbitmq/alert/alert.go",
    "groupTitle": "RabbitMQ_POST",
    "name": "TopicStock_alert"
  },
  {
    "type": "get",
    "url": "/api/v1/stock/:articleid",
    "title": "Request stock from an specific article.",
    "name": "GetStockFromAnArticle",
    "group": "Stock_Information",
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "articleid",
            "description": "<p>The unique id of an article.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "articleid",
            "description": "<p>Id of the article.</p>"
          },
          {
            "group": "200",
            "type": "Number",
            "optional": false,
            "field": "quantity",
            "description": "<p>Number of articles aviable.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success:",
          "content": "{\n    \"<articleid>\" : <articleid>,\n\t\t  \"quantity\" : <quantity>\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "ArticleNotFound:",
          "content": "\nHTTP/1.1 404 Not found\n{\"error\":\"Article not found\"}",
          "type": "json"
        }
      ]
    },
    "filename": "./main.go",
    "groupTitle": "Stock_Information"
  },
  {
    "type": "POST",
    "url": "/api/v1/stock/:articleid",
    "title": "Add stock to an article.",
    "name": "AddStockToAnArticle",
    "group": "Stock_Operations",
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "articleid",
            "description": "<p>The unique id of an article.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "quantity",
            "description": "<p>The quantity of an article.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success:",
          "content": "HTTP/1.1 200 Success\n\"Article added successfully\"",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "BadRequest:",
          "content": "HTTP/1.1 400 Bad Request\n\"Parameters needed\"",
          "type": "json"
        },
        {
          "title": "NotAutorized:",
          "content": "HTTP/1.1 401 Not Autorized\n\"Not Autorized\"",
          "type": "json"
        },
        {
          "title": "ArticleNotFound:",
          "content": "HTTP/1.1 404 Not Found\n\"Article not found\"",
          "type": "json"
        }
      ]
    },
    "filename": "./main.go",
    "groupTitle": "Stock_Operations"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./apidoc/main.js",
    "group": "_home_federico_GoWorkspace_src_github_com_fedegallar_stockmicroservice2018_apidoc_main_js",
    "groupTitle": "_home_federico_GoWorkspace_src_github_com_fedegallar_stockmicroservice2018_apidoc_main_js",
    "name": ""
  }
] });
