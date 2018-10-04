define({ "api": [
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
          "content": "{\n    \"5b97cd224a1aa37430480268\" : 50,\n}",
          "type": "json"
        },
        {
          "title": "Success-No-Content:",
          "content": "{\n    \"Code\" : 204,\n\t  \"Description\" : \"No content\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./main.go",
    "groupTitle": "Stock_Information"
  },
  {
    "type": "POST",
    "url": "/api/v1/stock",
    "title": "Add stock to an article. If there is no article, it will create a new one.",
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
          "content": "{\n    \"5b97cd224a1aa37430480268\" : 50,\n}",
          "type": "json"
        },
        {
          "title": "Success-No-Content:",
          "content": "{\n    \"Code\" : 204,\n\t  \"Description\" : \"No content\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./main.go",
    "groupTitle": "Stock_Operations"
  },
  {
    "type": "DELETE",
    "url": "/api/v1/stock/articleid",
    "title": "Remove stock from an article.",
    "name": "RemoveStockFromArticle",
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
          "content": "{\n    \"5b97cd224a1aa37430480268\" : 50,\n}",
          "type": "json"
        },
        {
          "title": "Success-No-Content:",
          "content": "{\n    \"Code\" : 204,\n\t  \"Description\" : \"No content\"\n}",
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
