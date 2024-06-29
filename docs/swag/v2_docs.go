// Package swag Code generated by swaggo/swag. DO NOT EDIT
package swag

import "github.com/swaggo/swag"

const docTemplatev2 = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "-V1 API",
            "url": "http://localhost:8081/swagger/v1/index.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v2/books": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Get a list of books in the the store",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/api.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "serializer.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfov2 holds exported Swagger Info so clients can modify it
var SwaggerInfov2 = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Test API",
	Description:      "Test API v2.0",
	InfoInstanceName: "v2",
	SwaggerTemplate:  docTemplatev2,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov2.InstanceName(), SwaggerInfov2)
}
