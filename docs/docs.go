// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Maintainer",
            "url": "https://github.com/Sveboo/url_shortener",
            "email": "svebo3348@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/Sveboo/url_shortener/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "post": {
                "description": "Shorten url provided in body and save it to storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shorten url",
                "parameters": [
                    {
                        "description": "Original url with protocol included",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Url shortened successfully",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Json is invalid",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    },
                    "422": {
                        "description": "Key 'url' is invalid or not provided",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    },
                    "500": {
                        "description": "Short url creation caused error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    }
                }
            }
        },
        "/{hash}": {
            "get": {
                "description": "Returns origin url by short form",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get original url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short url hash",
                        "name": "short_url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Short url exists in storage",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Short url not found in storage",
                        "schema": {
                            "$ref": "#/definitions/httpserver.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpserver.UserResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "some error message"
                },
                "url": {
                    "type": "string",
                    "example": "http://example.com"
                }
            }
        },
        "models.UserRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string",
                    "example": "http://example.com"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Url shortener documentation",
	Description:      "A collection of endpoints available to communicate with url shortener",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
