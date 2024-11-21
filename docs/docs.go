// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.example.com/support",
            "email": "support@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/otp-code": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Для ввода кода подтверждения, который приходит на почту",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Код подтверждения",
                "parameters": [
                    {
                        "description": "OTPRequest data",
                        "name": "OTPRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.OTPRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/registration": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Для регистрации ребуется передать уникальный username, пароль и уникальный email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users",
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "registrationData data",
                        "name": "registrationData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.registrationData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/createMarket": {
            "post": {
                "description": "Данный запрос позволяет создать магазин, если формат данных соответствует структуре models.Market",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "markets"
                ],
                "summary": "Создание магазина",
                "parameters": [
                    {
                        "description": "Market data",
                        "name": "market",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Market"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/createProduct": {
            "post": {
                "description": "Данный запрос позволяет создать продукт, если формат данных соответствует структуре models.Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Создание продукта",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "market",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/markets": {
            "get": {
                "description": "Данный запрос позволяет получить список всех магазинов, их данных и товаров внутри них",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "markets"
                ],
                "summary": "Все магазины",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Данный запрос позволяет получить список всех продуктов и данных о них",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Все продукты",
                "responses": {}
            }
        },
        "/products/{market_id}": {
            "get": {
                "description": "Данный запрос позволяет получить список товаров магазина по id магазина",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products",
                    "markets"
                ],
                "summary": "Продукты по marketID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Market ID",
                        "name": "market_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users": {
            "get": {
                "description": "Данный запрос позволяет получить список всех пользователей и их данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Все пользователи",
                "responses": {}
            }
        }
    },
    "definitions": {
        "handlers.OTPRequest": {
            "type": "object",
            "required": [
                "otp"
            ],
            "properties": {
                "otp": {
                    "type": "string"
                }
            }
        },
        "handlers.registrationData": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Market": {
            "type": "object"
        },
        "models.Product": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "My API",
	Description:      "This is a sample server for My API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
