// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/todo": {
            "get": {
                "description": "Получить список элементов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TodoItems"
                ],
                "summary": "Получить список элементов",
                "responses": {}
            },
            "post": {
                "description": "Создать элемент",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TodoItems"
                ],
                "summary": "Создать элемент",
                "parameters": [
                    {
                        "description": "Данные для создания нового элемента",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TodoItem"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/todo/{id}": {
            "get": {
                "description": "Получить элемент по id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TodoItems"
                ],
                "summary": "Получить элемент по id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TodoItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Удалить элемент",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TodoItems"
                ],
                "summary": "Удалить элемент",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TodoItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.TodoItem": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Список заданий",
	Description:      "Тестовое приложение для списка заданий",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
