// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
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
        "/tasks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.TaskResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error occured",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create new task",
                "parameters": [
                    {
                        "description": "Task request. Due field must be UTC time presented in format: yyyy-MM-ddThh:mm:ss",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Received bad request",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    },
                    "500": {
                        "description": "Internal error occured",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get task by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.TaskResponse"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    },
                    "500": {
                        "description": "Internal error occured",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update specified task by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task body. Due field must be UTC time presented in format: yyyy-MM-ddThh:mm:ss",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Received bad request",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    },
                    "500": {
                        "description": "Internal error occured",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    },
                    "500": {
                        "description": "Internal error occured",
                        "schema": {
                            "$ref": "#/definitions/middleware.customError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "middleware.customError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "request.CreateTaskRequest": {
            "type": "object",
            "required": [
                "description",
                "due",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "due": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.UpdateTaskRequest": {
            "type": "object",
            "required": [
                "description",
                "due",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "due": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.TaskResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
